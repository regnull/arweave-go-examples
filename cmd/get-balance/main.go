package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/everFinance/goar"
)

type CmdArgs struct {
	ArweaveURL string
	Address    string
}

func main() {
	var args CmdArgs
	flag.StringVar(&args.ArweaveURL, "arweave-url", "https://arweave.net", "arweave url")
	flag.StringVar(&args.Address, "address", "", "wallet address")
	flag.Parse()

	if args.Address == "" {
		log.Fatal("--address is required")
	}

	arClient := goar.NewClient(args.ArweaveURL)
	balance, err := arClient.GetWalletBalance(args.Address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%f\n", balance)
}
