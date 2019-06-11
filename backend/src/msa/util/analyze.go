package util

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"msa/types"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const CAPINFOS_TIMEFORMAT = "2006-01-02 15:04:05.000000"

// Analyzes all files from the config with `capinfos` command.
// Returns a config with additional information added to each of the entries.
func Analyze(config types.Config) types.Config {
	for i, tf := range config.TraceFiles {
		newTf, err := AnalyzeTraceFile(tf)
		if err != nil {
			log.Println(err)
			if !tf.AttackTrace {
				// only delete noise that could not be analyzed. Do not delete attack trace files when they error during analysis.
				delete(config.TraceFiles, i)
			}
			continue
		}
		config.TraceFiles[i] = newTf
	}
	return config
}

func AnalyzeTraceFile(tf types.TraceFile) (types.TraceFile, error) {
	// rTM: no header, table like tab separator, machine readable (long values)
	// aecsxy: start time, end time, number of packets, file size, average packet rate, average data rate
	cmd := exec.Command("capinfos", "-rTMaxyces", tf.Path)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	capInfosRes := strings.Split(out.String(), "\t")
	if err != nil && len(capInfosRes) != 7 {
		return tf, errors.New(fmt.Sprintf("Error running `capinfos` on %s: %v", tf.Path, err))
	}

	cmd = exec.Command("tcpdump", "-nn", "-r", tf.Path, "-qt", "ip")
	out.Reset()
	cmd.Stdout = &out
	err = cmd.Run()
	counts := make(map[string]int)
	tcpDumpRes := strings.Split(out.String(), "\n")
	for _, line := range tcpDumpRes {
		parts := strings.Split(line, " > ")
		if (len(parts) != 2) {
			continue
		}
		sa := strings.Split(parts[0], " ")
		ipPortA := sa[len(sa)-1]
		sb := strings.Split(parts[1], " ")
		ipPortB := sb[0]
		a := strings.Split(ipPortA, ".")
		b := strings.Split(ipPortB, ".")

		counts[strings.Join(a[0:4], ".")]++
		counts[strings.Join(b[0:4], ".")]++
	}
	max := 0
	for ip, count := range counts {
		if count > max {
			max = count
			tf.MostFrequentIp = ip
		}
	}

	if err != nil && tf.MostFrequentIp == "" {
		return tf, errors.New(fmt.Sprintf("Error running `tcpdump` on %s: %v", tf.Path, err))
	}

	// order: File name	Number of packets	File size (bytes)	Start time	End time	Data byte rate (bytes/sec)	Average packet rate (packets/sec)
	tf.Packets, _ = strconv.ParseInt(strings.TrimSpace(capInfosRes[1]), 10, 64)
	tf.FirstPacket, _ = time.Parse(CAPINFOS_TIMEFORMAT, capInfosRes[3])
	tf.LastPacket, _ = time.Parse(CAPINFOS_TIMEFORMAT, capInfosRes[4])
	tf.BytesPerSecond, _ = strconv.ParseFloat(strings.TrimSpace(capInfosRes[5]), 64)
	tf.PacketsPerSecond, _ = strconv.ParseFloat(strings.TrimSpace(capInfosRes[6]), 64)
	return tf, nil
}
