package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"msa/util"
	"path"
)

var (
	replaceCmd = &cobra.Command{
		Use:   "replace",
		Short: "replace ip addresses in pcaps",
		Long:  `This command replaces IP addresses in the configured pcaps and outputs new pcaps`,
		Run:   replace,
	}
	replacements string
	id           uint32
)

func init() {
	replaceCmd.Flags().StringVarP(&replacements, "replace", "r", "", "comma delimited string of IP pairs to replace (orig:replace,...)")
	replaceCmd.MarkFlagRequired("replace")

	replaceCmd.Flags().Uint32VarP(&id, "id", "i", 0, "id of trace to work with. use `list` command to find the IDs")
}

func replace(cobraCmd *cobra.Command, args []string) {
	tf, ok := Config.TraceFiles[id]
	if !ok {
		log.Fatal("Unknown file id.")
	}

	outFile := path.Join(Config.OutPath, path.Base(tf.Path)+"_ips_replaced")
	tfReplaced, err := util.Replace(replacements, tf.Path, outFile)
	if err != nil {
		log.Fatal(err)
	}
	tfReplaced, err = util.AnalyzeTraceFile(tfReplaced)
	if err != nil {
		log.Fatal(err)
	}
	if exportFile != "" {
		Config.TraceFiles[tfReplaced.Id] = tfReplaced
		err = Config.WriteToFile(exportFile)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Export finished")
	}
}
