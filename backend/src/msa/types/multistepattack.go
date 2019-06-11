package types

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type TimeLineEntry struct {
	Type string `json:"type"`
	Id   uint32 `json:"id"`
}

type TimeLineDay = []TimeLineEntry

type Replacement struct {
	IpA string `json:"ipA"`
	IpB string `json:"ipB"`
}

type MultiStepAttackRequest struct {
	TimeLine     []TimeLineDay
	Replacements []Replacement `json:"replacements"`
	Name         string        `json:"name"`
}

type MultiStepAttack struct {
	Attacks    Attacks       `json:"attacks"`
	TraceFiles TraceFiles    `json:"traceFiles"`
	TimeLine   []TimeLineDay `json:"timeline"`
	Name       string        `json:"name"`
}

type MultiStepAttacks = map[string]MultiStepAttack

func (msa *MultiStepAttack) WriteToFile(outPath string) error {
	jsonConfig, err := json.Marshal(msa)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(outPath, msa.Name)+".json", jsonConfig, 0644)
}

func (repl *Replacement) ToReplacementString() string {
	return fmt.Sprintf("[%s]:[%s]", repl.IpA, repl.IpB)
}

// loads a json file that contains state for a MSA
func LoadMultiStepAttack(filePath string) (msa MultiStepAttack, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&msa)
	return
}
