package types

import (
	"encoding/json"
	"os"
	"time"
)

type attack struct {
	Attackers []string  `json:"attackers"`
	Victims   []string  `json:"victims"`
	Name      string    `json:"name"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
}

type groundTruth struct {
	Files   []string `json:"files"`
	Attacks []attack `json:"attacks"`
}

type configFile struct {
	GroundTruth         []groundTruth `json:"groundTruth"`
	UnclassifiedTraffic []string      `json:"unclassifiedTraffic"`
	OutPath             string        `json:"outPath"`
}

// Parses a JSON config from hard disk. Returns an error when JSON is malformed
func readConfigurationFile(path string) (cf configFile, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&cf)
	return
}

// Takes the file-struct and expands it into an internal config struct
func mapConfigFile(cf configFile) (c Config) {
	c.Attacks = make(map[uint32]Attack)
	c.TraceFiles = make(map[uint32]TraceFile)
	for _, gt := range cf.GroundTruth {
		var traceIds []uint32
		for _, path := range gt.Files {
			for _, tf := range TraceFilesFromPath(path) {
				traceIds = append(traceIds, tf.Id)
				tf.AttackTrace = true
				c.TraceFiles[tf.Id] = tf
			}
		}
		for _, atk := range gt.Attacks {
			var attack Attack
			attack.Attackers = atk.Attackers
			attack.Victims = atk.Victims
			attack.Name = atk.Name
			attack.Start = atk.Start
			attack.End = atk.End
			attack.Traces = traceIds
			attack.Id = attack.Hash()
			c.Attacks[attack.Id] = attack
		}
	}
	for _, path := range cf.UnclassifiedTraffic {
		for _, tf := range TraceFilesFromPath(path) {
			c.TraceFiles[tf.Id] = tf
		}
	}
	c.OutPath = cf.OutPath
	return
}

func LoadConfigurationFile(file string) (c Config, err error) {
	cf, err := readConfigurationFile(file)
	if err == nil {
		c = mapConfigFile(cf)
	}
	return
}
