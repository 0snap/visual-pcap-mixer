package util

import (
	"errors"
	"fmt"
	"log"
	"msa/types"
	"os/exec"
	"path"
	"strings"
)

func Replace(replacements, inFile, outFile string) (tfNew types.TraceFile, err error) {
	if replacements == "" {
		return tfNew, errors.New("Empty replacements string")
	}

	cmd := exec.Command("tcprewrite", "--infile="+inFile, "--outfile="+outFile, "--srcipmap="+replacements, "--dstipmap="+replacements)
	log.Printf("Going to execute:       %v", cmd.Args)
	err = cmd.Run()
	if err != nil {
		return
	}
	tfNew.Path = outFile

	tfNew, err = AnalyzeTraceFile(tfNew)
	if err != nil {
		log.Printf("Error analyzing the new file!")
		return
	}
	tfNew.Id = tfNew.Hash()
	return
}

func ReplaceInFile(replacements []types.Replacement, tf types.TraceFile, outPath string) (tfNew types.TraceFile, err error) {
	var repls []string
	for _, repl := range replacements {
		repls = append(repls, fmt.Sprintf("[%s]:[%s]", repl.IpA, repl.IpB))
	}

	outFile := path.Join(outPath, path.Base(tf.Path)+"_ips_replaced")

	return Replace(strings.Join(repls, ","), tf.Path, outFile)
}

func ReplaceInAllFiles(replacements []types.Replacement, tfs []types.TraceFile, outPath string) (tfsNew []types.TraceFile, err error) {
	for _, tf := range tfs {
		tfNew, err := ReplaceInFile(replacements, tf, outPath)
		if err != nil {
			return tfsNew, err
		}
		tfsNew = append(tfsNew, tfNew)
	}
	return
}

func ReplaceInAttack(replacements []types.Replacement, atk types.Attack, allTraceFiles types.TraceFiles, outPath string) (atkNew types.Attack, tfsNew []types.TraceFile, err error) {
	atkNew.Name = atk.Name
	atkNew.Start = atk.Start
	atkNew.End = atk.End

	for _, attacker := range atk.Attackers {
		for _, repl := range replacements {
			if attacker == repl.IpA {
				attacker = repl.IpB
			} else if attacker == repl.IpB {
				attacker = repl.IpA
			}
		}
		atkNew.Attackers = append(atkNew.Attackers, attacker)
	}
	for _, victim := range atk.Victims {
		for _, repl := range replacements {
			if victim == repl.IpA {
				victim = repl.IpB
			} else if victim == repl.IpB {
				victim = repl.IpA
			}
		}
		atkNew.Victims = append(atkNew.Victims, victim)
	}

	var tfs []types.TraceFile
	for _, traceId := range atk.Traces {
		tfs = append(tfs, allTraceFiles[traceId])
	}
	tfsNew, err = ReplaceInAllFiles(replacements, tfs, outPath)
	if err != nil {
		return
	}
	for _, tfNew := range tfsNew {
		atkNew.Traces = append(atkNew.Traces, tfNew.Id)
	}
	atkNew.Id = atkNew.Hash()
	return
}
