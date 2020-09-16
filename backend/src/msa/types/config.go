package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

type TraceFile struct {
	Path             string    `json:"path"`
	Packets          int64     `json:"packets"`
	FirstPacket      time.Time `json:"firstPacket"`
	LastPacket       time.Time `json:"lastPacket"`
	PacketsPerSecond float64   `json:"packetsPerSecond"`
	BytesPerSecond   float64   `json:"bytesPerSecond"`
	Id               uint32    `json:"id"`
	AttackTrace      bool      `json:"attackTrace"`
	MostFrequentIp   string    `json:"mostFrequentIp"`
}

type Attack struct {
	Attackers []string  `json:"attackers"`
	Victims   []string  `json:"victims"`
	Name      string    `json:"name"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Traces    []uint32  `json:"traces"`
	Id        uint32    `json:"id"`
}

type TraceFiles = map[uint32]TraceFile
type Attacks = map[uint32]Attack

type Config struct {
	Attacks          Attacks          `json:"attacks"`
	TraceFiles       TraceFiles       `json:"traceFiles"`
	OutPath          string           `json:"outPath"`
	MultiStepAttacks MultiStepAttacks `json:"multistepattacks"`
}

func TraceFilesFromPath(filePath string) (result []TraceFile) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Printf("Error. Cannot access path at '%s'\n", filePath)
		log.Println(err)
	}
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(filePath)
		if err != nil {
			log.Println("Error. Cannot read directory '%s'\n", filePath)
			log.Println(err)
		}
		for _, f := range files {
			var tf TraceFile
			tf.Path = path.Join(filePath, f.Name())
			tf.Id = tf.Hash()
			result = append(result, tf)
		}
	} else {
		var tf TraceFile
		tf.Path = filePath
		tf.Id = tf.Hash()
		result = append(result, tf)
	}
	return
}

func (tf *TraceFile) ToString() string {
	return fmt.Sprintf("%-10d: %50s | %v -- %v | packets: %d | packets/sec: %f | bytes/sec: %f\n", tf.Id, tf.Path, tf.FirstPacket, tf.LastPacket, tf.Packets, tf.PacketsPerSecond, tf.BytesPerSecond)
}

func (tf *TraceFile) Hash() uint32 {
	h := fnv.New32a()
	h.Write([]byte(tf.Path))
	return h.Sum32()
}

func (atk *Attack) Hash() uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", atk)))
	return h.Sum32()
}

func (cfg *Config) WriteToFile(exportFile string) error {
	jsonConfig, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(exportFile, jsonConfig, 0644)
}

// loads a json file that was exported by this tool and contains the interna state (annotated pcaps with data from `capinfos`)
func LoadStateFile(path string) (c Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&c)
	return
}

func (cfg *Config) GetEntry(id uint32) interface{} {
	if entry, found := cfg.TraceFiles[id]; found {
		return entry
	}
	if entry, found := cfg.Attacks[id]; found {
		return entry
	}
	return nil
}

func (cfg *Config) AddMsa(msa MultiStepAttack) error {
	if _, exists := cfg.MultiStepAttacks[msa.Name]; exists {
		return errors.New("MultiStepAttack with this name already exists!")
	}

	cfg.MultiStepAttacks[msa.Name] = msa
	for id, atk := range msa.Attacks {
		cfg.Attacks[id] = atk
	}
	for id, tf := range msa.TraceFiles {
		cfg.TraceFiles[id] = tf
	}
	return nil
}
