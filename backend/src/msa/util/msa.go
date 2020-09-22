package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"msa/types"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

const MAX_WORKERS int = 50

func copyTraceFile(tf types.TraceFile, outPath string) (tfNew types.TraceFile, err error) {
	outFile := path.Join(outPath, path.Base(tf.Path))
	err = copyFile(tf.Path, outFile)
	if err != nil {
		return tfNew, err
	}
	tf.Path = outFile
	tf.Id = tf.Hash()
	return tf, nil
}

func copyFile(src, dst string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dst, data, 0644)
}

func removeFile(path, allowedPath string) {
	if strings.HasPrefix(path, allowedPath) {
		os.Remove(path)
	}
}

type Item struct {
	Date  time.Time
	Day   int
	Index int
	Path  string
	Entry types.TimeLineEntry
}

type TraceResult struct {
	Day   int
	Index int
	Trace types.TraceFile
}

type AttackResult struct {
	Day    int
	Index  int
	Attack types.Attack
	Traces []types.TraceFile
}

func processTrace(tf *types.TraceFile, item Item, repl *[]types.Replacement) TraceResult {
	// 1. Adjust timestamps
	tfNew, err := SetDate(item.Date, *tf, item.Path)
	if err != nil {
		log.Fatalf("Error during time adjustment for %v: %v", tf.Path, err)
	}

	if len(*repl) != 0 {
		// Save path of the trace to delete later
		tfPath := tfNew.Path

		// 2. Replace IPs
		tfNew, err = ReplaceInFile(*repl, tfNew, item.Path)
		if err != nil {
			log.Fatalf("Error during IP replacement for %v: %v", tfNew.Path, err)
		}

		// Remove old file
		removeFile(tfPath, item.Path)
	}

	return TraceResult{item.Day, item.Index, tfNew}
}

func processAttack(atk *types.Attack, traces *map[uint32]types.TraceFile, item Item, repl *[]types.Replacement) AttackResult {
	// 1. Adjust timestamps
	atkNew, tfs, err := SetAttackDate(item.Date, *atk, *traces, item.Path)
	if err != nil {
		log.Fatalf("Error during time adjustment for %v: %v", atk.Name, err)
	}

	if len(*repl) != 0 {
		// Save paths of the traces to delete later
		tfPaths := make([]string, 0, len(tfs))
		for _, tf := range tfs {
			tfPaths = append(tfPaths, tf.Path)
		}

		// 2. Replace IPs
		atkNew, tfs, err = ReplaceInAttack(*repl, atkNew, tfs, item.Path)
		if err != nil {
			log.Fatalf("Error during IP replacement for %v: %v", atk.Name, err)
		}

		// Remove old files
		for _, p := range tfPaths {
			removeFile(p, item.Path)
		}
	}

	return AttackResult{item.Day, item.Index, atkNew, tfs}
}

func processItem(item Item, traces *map[uint32]types.TraceFile, attacks *map[uint32]types.Attack, repl *[]types.Replacement, tfRes chan<- TraceResult, atkRes chan<- AttackResult) {
	if item.Entry.Type == "traceFile" {
		tf := (*traces)[item.Entry.Id]
		res := processTrace(&tf, item, repl)
		tfRes <- res
	} else if item.Entry.Type == "attack" {
		atk := (*attacks)[item.Entry.Id]
		res := processAttack(&atk, traces, item, repl)
		atkRes <- res
	}
}

// FIXME: This effectively crashes on error and does not properly use `err`. Might want to revisit that
func GenerateMultiStepAttack(req types.MultiStepAttackRequest, config types.Config) (msa types.MultiStepAttack, err error) {
	parentOutPath := path.Join(config.OutPath, req.Name)
	os.MkdirAll(parentOutPath, os.ModePerm)

	msa.Name = req.Name
	msa.TraceFiles = make(types.TraceFiles)
	msa.Attacks = make(types.Attacks)

	msa.TimeLine = make([]types.TimeLineDay, len(req.TimeLine))
	pathsPerDay := make([]string, len(req.TimeLine))

	for i := range msa.TimeLine {
		msa.TimeLine[i] = make([]types.TimeLineEntry, len(req.TimeLine[i]))

		pathsPerDay[i] = path.Join(parentOutPath, fmt.Sprintf("day-%d", i+1))
		os.MkdirAll(pathsPerDay[i], os.ModePerm)
	}

	firstDay := req.TimeLine[0]
	firstEntry := firstDay[0]
	prevTraceFile := config.TraceFiles[firstEntry.Id]
	if firstEntry.Type == "attack" {
		atk := config.Attacks[firstEntry.Id]
		lastId := atk.Traces[len(atk.Traces)-1]
		prevTraceFile = config.TraceFiles[lastId]
	}

	startDate := prevTraceFile.FirstPacket

	// Channels
	items := make(chan Item, 10240)
	tfResults := make(chan TraceResult, 64)
	// FIXME: If we have more than 64 attacks, the loops below don't terminate..
	atkResults := make(chan AttackResult, 64)
	// WaitGroups
	var wgIn sync.WaitGroup
	var wgOut sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < MAX_WORKERS; i++ {
		go func() {
			// Receive item to process
			for item := range items {
				wgIn.Done()
				processItem(item, &config.TraceFiles, &config.Attacks, &req.Replacements, tfResults, atkResults)
				wgOut.Done()
			}
		}()
	}

	// Enqueue all entries across all days
	for day, entries := range req.TimeLine {
		outPath := pathsPerDay[day]
		date := startDate.Add(time.Hour * 24 * time.Duration(day))

		// Add number of entries to WaitGroups
		numEntries := len(entries)
		wgIn.Add(numEntries)
		wgOut.Add(numEntries)

		// Enqueue entries as items
		for idx, entry := range entries {
			items <- Item{date, day, idx, outPath, entry}
		}
	}

	// Wait for all input items to be consumed and close input channel
	go func() {
		// This terminates idle workers once all items have been retrieved
		// It also has to be in a goroutine to not block the main thread
		wgIn.Wait()
		close(items)
	}()
	// Wait for all workers to finish and close output channels
	go func() {
		// This is required as the range loops below don't exit otherwise
		// It also has to be in a goroutine to not block the main thread
		wgOut.Wait()
		close(tfResults)
		close(atkResults)
	}()

	// Collect results and update resulting MSA
	for result := range tfResults {
		tf := result.Trace

		msa.TraceFiles[tf.Id] = tf
		msa.TimeLine[result.Day][result.Index] = types.TimeLineEntry{"traceFile", tf.Id}
	}

	for result := range atkResults {
		atk := result.Attack

		msa.Attacks[atk.Id] = atk
		for _, tf := range result.Traces {
			msa.TraceFiles[tf.Id] = tf
		}
		msa.TimeLine[result.Day][result.Index] = types.TimeLineEntry{"attack", atk.Id}
	}

	return msa, nil
}
