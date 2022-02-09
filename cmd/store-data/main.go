package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
)

type MapFlags map[string]string

func (i *MapFlags) String() string {
	return "my string representation"
}

func (i *MapFlags) Set(value string) error {
	parts := strings.Split(value, "=")
	if len(parts) != 2 {
		return errors.New("invalid value")
	}
	(*i)[parts[0]] = parts[1]
	return nil
}

type CmdArgs struct {
	KeyFile    string
	ArweaveURL string
	DataFile   string
	Tags       MapFlags
}

func main() {
	var args CmdArgs
	args.Tags = make(map[string]string)
	flag.StringVar(&args.KeyFile, "key-file", "", "key file")
	flag.StringVar(&args.ArweaveURL, "arweave-url", "https://arweave.net", "arweave url")
	flag.StringVar(&args.DataFile, "data-file", "", "data file")
	flag.Var(&args.Tags, "tag", "tags")
	flag.Parse()

	if args.KeyFile == "" {
		log.Fatal("--key-file is required")
	}

	if args.DataFile == "" {
		log.Fatal("--data-file is required")
	}

	data, err := ioutil.ReadFile(args.DataFile)
	if err != nil {
		log.Fatal(err)
	}

	wallet, err := goar.NewWalletFromPath(args.KeyFile, args.ArweaveURL)
	if err != nil {
		log.Fatal(err)
	}

	var tags []types.Tag
	for k, v := range args.Tags {
		tags = append(tags, types.Tag{
			Name:  k,
			Value: v,
		})
	}
	id, err := wallet.SendData(data, tags)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("id: %s\n", id)
}
