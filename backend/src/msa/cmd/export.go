package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"msa/util"
)

var (
	exportCmd = &cobra.Command{
		Use:   "export",
		Short: "analyzes the user config and exports the internal state",
		Long:  `the internal state export can be read again for later runs. it contains all processed information about the analyzed pcaps.`,
		Run:   export,
	}
)

func export(cmd *cobra.Command, args []string) {

	if exportFile == "" {
		exportFile = "./export.json"
	}

	if !StateFileInitialized {
		log.Println("Starting analysis ...")
		Config = util.Analyze(Config)
	}

	Config.WriteToFile(exportFile)
	log.Println("Export finished.")
}
