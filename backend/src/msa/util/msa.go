package util

import (
	"io/ioutil"
	"msa/types"
	"os"
	"path"
	"strings"
	"time"
)

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

func GenerateMultiStepAttack(msar types.MultiStepAttackRequest, config types.Config) (msa types.MultiStepAttack, err error) {
	outPath := path.Join(config.OutPath, msar.Name)
	os.MkdirAll(outPath, os.ModePerm)

	msa.Name = msar.Name
	msa.TraceFiles = make(types.TraceFiles)
	msa.Attacks = make(types.Attacks)

	msa.TimeLine = make([]types.TimeLineDay, len(msar.TimeLine))
	for i := range msa.TimeLine {
		msa.TimeLine[i] = make([]types.TimeLineEntry, len(msar.TimeLine[i]))
	}

	var tf types.TraceFile

	firstDay := msar.TimeLine[0]
	firstEntry := firstDay[0]
	prevTraceFile := config.TraceFiles[firstEntry.Id]
	if firstEntry.Type == "attack" {
		atk := config.Attacks[firstEntry.Id]
		lastId := atk.Traces[len(atk.Traces)-1]
		prevTraceFile = config.TraceFiles[lastId]
	}

	currentDate := prevTraceFile.FirstPacket
	// connection of timestamps in the PCAPs
	for i, day := range msar.TimeLine {
		for e, entry := range day {
			if entry.Type == "traceFile" {
				tf, err = SetDate(currentDate, config.TraceFiles[entry.Id], outPath)
				if err != nil {
					return
				}
				msa.TraceFiles[tf.Id] = tf
				msa.TimeLine[i][e] = types.TimeLineEntry{"traceFile", tf.Id}
			} else if entry.Type == "attack" {
				atk, tfs, err := SetAttackDate(currentDate, config.Attacks[entry.Id], config.TraceFiles, outPath)
				if err != nil {
					return msa, err
				}
				msa.Attacks[atk.Id] = atk
				for _, tf := range tfs {
					msa.TraceFiles[tf.Id] = tf
					prevTraceFile = tf
				}
				msa.TimeLine[i][e] = types.TimeLineEntry{"attack", atk.Id}
			}
		}
		currentDate = currentDate.Add(time.Hour * 24)
	}

	// IP address replacements
	if len(msar.Replacements) != 0 {
		for i, day := range msa.TimeLine {
			for e, entry := range day {
				if entry.Type == "traceFile" {
					tf, err = ReplaceInFile(msar.Replacements, msa.TraceFiles[entry.Id], outPath)
					if err != nil {
						return
					}
					removeFile(msa.TraceFiles[entry.Id].Path, outPath)
					delete(msa.TraceFiles, entry.Id)
					msa.TraceFiles[tf.Id] = tf
					msa.TimeLine[i][e] = types.TimeLineEntry{"traceFile", tf.Id}
				} else if entry.Type == "attack" {
					atk, tfs, err := ReplaceInAttack(msar.Replacements, msa.Attacks[entry.Id], msa.TraceFiles, outPath)
					if err != nil {
						return msa, err
					}
					msa.Attacks[atk.Id] = atk
					for _, tf := range tfs {
						msa.TraceFiles[tf.Id] = tf
					}
					for _, oldAtkTrace := range msa.Attacks[entry.Id].Traces {
						removeFile(msa.TraceFiles[oldAtkTrace].Path, outPath)
						delete(msa.TraceFiles, oldAtkTrace)
					}
					delete(msa.Attacks, entry.Id)
					msa.TimeLine[i][e] = types.TimeLineEntry{"attack", atk.Id}
				}
			}
		}
	}

	return msa, nil
}
