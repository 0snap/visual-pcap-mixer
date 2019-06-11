package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"msa/util"
	"time"
)

var (
	analyze bool
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "lists all pcaps",
		Long:  `This command lists all configured pcaps and some meta information`,
		Run:   list,
	}
)

func init() {
	listCmd.Flags().BoolVarP(&analyze, "analyze", "a", false, "uses `capinfos` binary to analyze pcaps. prints detailed information per file.")
}

func list(cmd *cobra.Command, args []string) {
	if analyze {
		Config = util.Analyze(Config)
	}

	fmt.Println("Known attacks and durations")
	for _, atk := range Config.Attacks {
		durString := fmt.Sprintf("%v  --  %v (%v)", atk.Start.Format(time.RFC3339), atk.End.Format(time.RFC3339), atk.End.Sub(atk.Start))
		fmt.Printf("%d: %20s | %-70s | %-30v --atk--> %-30v \n", atk.Id, atk.Name, durString, atk.Attackers, atk.Victims)
	}
	fmt.Println("\nTracefiles")
	for _, tf := range Config.TraceFiles {
		fmt.Printf("%s\n", tf.ToString())
	}

}
