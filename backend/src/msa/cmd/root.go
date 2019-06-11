package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"msa/types"
)

var (
	cfgFile    string
	stateFile  string
	exportFile string
	outPath    string

	Config               types.Config
	StateFileInitialized bool

	rootCmd = &cobra.Command{
		Use:   "msa",
		Short: "A generator for joining attack pcaps to multi-step-attacks",
		Long:  `This tool replaces IP addresses packet captures. It offers nice means of configuration`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initializeState)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./config.json", "path to config file (default is $(PWD)/config.json")
	rootCmd.PersistentFlags().StringVarP(&stateFile, "statefile", "s", "", "(optional) path to statefile. It contains an already processed internal config state.")
	rootCmd.PersistentFlags().StringVarP(&exportFile, "export", "e", "", "path to output json file")
	rootCmd.PersistentFlags().StringVarP(&outPath, "out", "o", "", "path to output any modified pcaps (default is $(PWD)/out)")

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(replaceCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(connectCmd)
}

func initializeState() {
	initConfig()           // reads the config / state file and stitches everything together to get internal state
	initMultiStepAttacks() // tries to read the "output" folder and find some already existing MSA
}

func initConfig() {
	var err error
	var c types.Config
	if stateFile != "" {
		c, err = types.LoadStateFile(stateFile)
		StateFileInitialized = true
		outPath = c.OutPath
	} else {
		// init without statefile (normal init with raw user-created config)
		c, err = types.LoadConfigurationFile(cfgFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	if outPath == "" {
		if c.OutPath != "" {
			outPath = c.OutPath
		} else {
			outPath = "./out"
		}
	}
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		os.Mkdir(outPath, os.ModePerm)
	}
	c.OutPath, err = filepath.Abs(outPath)
	if err != nil {
		log.Fatal(err)
	}
	if c.MultiStepAttacks == nil {
		c.MultiStepAttacks = make(types.MultiStepAttacks)
	}
	Config = c
	return
}

func initMultiStepAttacks() {
	files, err := ioutil.ReadDir(Config.OutPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.Mode().IsRegular() || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		msa, err := types.LoadMultiStepAttack(path.Join(Config.OutPath, file.Name()))
		if err != nil {
			log.Printf("Error loading MSA state file in output directory '%s': %v", file.Name(), err)
			continue
		}
		Config.AddMsa(msa)
	}
}
