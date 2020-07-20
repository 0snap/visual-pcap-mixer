package util

import (
	"errors"
	"fmt"
	"log"
	"msa/types"
	"os/exec"
	"path"
	"time"
)

const EMPTY_DATE = "0001-01-01 00:00:00 +0000"

func validateTraces(tfA, tfB types.TraceFile) error {
	emptyTime, _ := time.Parse("", EMPTY_DATE)
	// TODO: delay between first and last, according to packet rates
	if tfA.LastPacket == emptyTime || tfB.FirstPacket == emptyTime || tfA.Path == "" || tfB.Path == "" || tfA.Path == tfB.Path {
		return errors.New(fmt.Sprintf("Malformed tracefiles: %d - %d", tfA.Id, tfB.Id))
	}
	return nil
}

func validateTrace(tf types.TraceFile) error {
	emptyTime, _ := time.Parse("", EMPTY_DATE)
	if tf.FirstPacket == emptyTime || tf.LastPacket == emptyTime || tf.Path == "" {
		return errors.New(fmt.Sprintf("Malformed tracefile: %d", tf.Id))
	}
	return nil
}

// adjusts all timestamps in `tfB` such that both captures appear to be recorded right after another.
// the tracefiles must be analyzed
func ConnectTimes(tfA, tfB types.TraceFile, outPath string) (tfNew types.TraceFile, err error) {
	if err = validateTraces(tfA, tfB); err != nil {
		return
	}

	outFile := path.Join(outPath, path.Base(tfB.Path)+"_time-adjusted")

	timeDiff := -1 * tfB.FirstPacket.Sub(tfA.LastPacket)
	return applyTimeDiff(timeDiff, tfB, outFile)
}

func ConnectAttackTimes(tfA types.TraceFile, atk types.Attack, traceFiles types.TraceFiles, outPath string) (atkNew types.Attack, attackTraces []types.TraceFile, err error) {
	atkNew.Attackers = atk.Attackers
	atkNew.Victims = atk.Victims
	atkNew.Name = atk.Name

	// adjust all traces in the attack with the same offset (first trace of attack to previos tracefile),
	// such that this does not mess with the timings across the attack traces but keeps the distribution.
	timeDiff := -1 * traceFiles[atk.Traces[0]].FirstPacket.Sub(tfA.LastPacket)

	atkNew.Start = atk.Start.Add(timeDiff)
	atkNew.End = atk.End.Add(timeDiff)

	for _, tfId := range atk.Traces {
		tfB := traceFiles[tfId]
		if err = validateTraces(tfA, tfB); err != nil {
			return
		}

		outFile := path.Join(outPath, path.Base(tfB.Path)+"_time-adjusted")

		var tfNew types.TraceFile
		tfNew, err = applyTimeDiff(timeDiff, tfB, outFile)
		if err != nil {
			return
		}

		attackTraces = append(attackTraces, tfNew)
		atkNew.Traces = append(atkNew.Traces, tfNew.Id)
	}
	atkNew.Id = atkNew.Hash()
	return
}

func applyTimeDiff(timeDiff time.Duration, tf types.TraceFile, outFile string) (tfNew types.TraceFile, err error) {
	log.Printf("exec: editcap -F pcap -t %v, %v %v", fmt.Sprintf("%f", timeDiff.Seconds()), tf.Path, outFile)
	cmd := exec.Command("editcap", "-F", "pcap", "-t", fmt.Sprintf("%f", timeDiff.Seconds()), tf.Path, outFile)
	err = cmd.Run()
	if err != nil {
		err = errors.New(fmt.Sprintf("Error running `editcap` on %s: %v\n", tf.Path, err))
		return
	}
	tfNew.Path = outFile
	tfNew, err = AnalyzeTraceFile(tfNew)
	if err != nil {
		log.Printf("Error analyzing the new connected file. Cannot add to config!")
		return
	}
	tfNew.Id = tfNew.Hash()
	return
}

// adjusts all timestamps in `tf` such that the capture appears to be recorded on the passed date (day)
func SetDate(date time.Time, tf types.TraceFile, outPath string) (tfNew types.TraceFile, err error) {
	if err = validateTrace(tf); err != nil {
		return
	}

	outFile := path.Join(outPath, path.Base(tf.Path)+"_time-adjusted")

	hours := int(tf.FirstPacket.Sub(date).Hours())
	duration, _ := time.ParseDuration(fmt.Sprintf("%dh", hours))
	return applyTimeDiff(-1*duration, tf, outFile)
}

// Sets the date (day) for all TFs of the attack, without modifying the internal timings of the tracefiles.
func SetAttackDate(date time.Time, atk types.Attack, traceFiles types.TraceFiles, outPath string) (atkNew types.Attack, attackTraces []types.TraceFile, err error) {
	atkNew.Attackers = atk.Attackers
	atkNew.Victims = atk.Victims
	atkNew.Name = atk.Name

	// adjust all traces in the attack with the same offset (first trace of attack to previos tracefile),
	// such that this does not mess with the timings across the attack traces but keeps the distribution.
	hours := int(traceFiles[atk.Traces[0]].FirstPacket.Sub(date).Hours())
	duration, _ := time.ParseDuration(fmt.Sprintf("%dh", hours))
	timeDiff := -1 * duration

	atkNew.Start = atk.Start.Add(timeDiff)
	atkNew.End = atk.End.Add(timeDiff)

	for _, tfId := range atk.Traces {
		tfB := traceFiles[tfId]
		if err = validateTrace(tfB); err != nil {
			return
		}

		outFile := path.Join(outPath, path.Base(tfB.Path)+"_time-adjusted")

		var tfNew types.TraceFile
		tfNew, err = applyTimeDiff(timeDiff, tfB, outFile)
		if err != nil {
			return
		}

		attackTraces = append(attackTraces, tfNew)
		atkNew.Traces = append(atkNew.Traces, tfNew.Id)
	}
	atkNew.Id = atkNew.Hash()
	return
}
