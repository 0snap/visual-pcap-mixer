package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"msa/util"
)

var (
	connectCmd = &cobra.Command{
		Use:   "connect",
		Short: "connect timestamps in pcaps",
		Long:  `This command synchronizes the timestamps for two of the configured pcaps and outputs new pcaps that looks as if it was recorded right after the first one`,
		Run:   connect,
	}
	idA uint32
	idB uint32
)

func init() {
	connectCmd.Flags().Uint32VarP(&idA, "idA", "a", 0, "the id of the first file (use `list` to see the IDs)")
	connectCmd.Flags().Uint32VarP(&idB, "idB", "b", 0, "the id of the second file (use `list` to see the IDs)")
	connectCmd.MarkFlagRequired("idA")
	connectCmd.MarkFlagRequired("idB")
}

func connect(cobraCmd *cobra.Command, args []string) {

	if !StateFileInitialized {
		log.Printf("Starting analysis ...")
		Config = util.Analyze(Config)
	}

	tfA, ok := Config.TraceFiles[idA]
	if !ok {
		log.Fatal("Unknown file `idA`.")
	}
	tfB, ok := Config.TraceFiles[idB]
	if !ok {
		log.Fatal("Unknown file `idB`.")
	}

	tfNewB, err := util.ConnectTimes(tfA, tfB, Config.OutPath)
	if err != nil {
		log.Fatal(err)
	}
	tfNewB, err = util.AnalyzeTraceFile(tfNewB)
	if err != nil {
		log.Fatal(err)
	}
	if exportFile != "" {
		Config.TraceFiles[tfNewB.Id] = tfNewB
		err = Config.WriteToFile(exportFile)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Export finished")
	}
}
