package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"

	"msa/types"
	"msa/util"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "start a server",
		Long:  `This command starts a server that will host all configured attacks and accept commands to manipulate the pcaps`,
		Run:   server,
	}
	ip   string
	port int
)

func init() {
	serverCmd.Flags().StringVarP(&ip, "listen", "l", "127.0.0.1", "IP to use for listening, default 127.0.0.1")
	serverCmd.Flags().IntVarP(&port, "port", "p", 1337, "port to listen on, default 1337")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func server(cmd *cobra.Command, args []string) {

	if !StateFileInitialized {
		fmt.Println("Server now analyzes all configured pcaps ...")
		Config = util.Analyze(Config)
	}

	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/msa", msaHandler)
	log.Printf("Starting server at %s:%d\n", ip, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", ip, port), nil))
}

// --------------- there is a handler for each CLI command -----------------

func listHandler(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(Config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func msaHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	decoder := json.NewDecoder(r.Body)
	var msar types.MultiStepAttackRequest
	err := decoder.Decode(&msar)
	if err != nil {
		http.Error(w, "Received malformed `MultiStepAttack request`, unable to parse.", http.StatusBadRequest)
		return
	}
	if _, exists := Config.MultiStepAttacks[msar.Name]; exists {
		http.Error(w, fmt.Sprint("A multistep attack with name already exists."), http.StatusBadRequest)
		return
	}
	msa, err := util.GenerateMultiStepAttack(msar, Config)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error applying the multi-step scenario: %v", err), http.StatusInternalServerError)
		return
	}
	msa.WriteToFile(Config.OutPath)
	Config.AddMsa(msa)
	js, err := json.Marshal(msa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
