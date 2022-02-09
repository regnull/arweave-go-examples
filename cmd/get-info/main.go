package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/everFinance/goar"
)

type CmdArgs struct {
	ArweaveURL string
}

func main() {
	var args CmdArgs
	flag.StringVar(&args.ArweaveURL, "arweave-url", "https://arweave.net", "arweave url")
	flag.Parse()

	arClient := goar.NewClient(args.ArweaveURL)
	info, err := arClient.GetInfo()
	if err != nil {
		log.Fatal(err)
	}

	// Format info nicely.
	jsonInfo, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(jsonInfo))
}
