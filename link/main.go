package main

import (
	"errors"
	"flag"
	"link/linklib"
	"log"
	"os"
)

type options struct {
	inputFile  string
	outputFile string
}

func getOptions(args []string) (options, error) {
	var opts options

	fs := flag.NewFlagSet("link", flag.ExitOnError)
	fs.StringVar(&opts.inputFile, "inputFile", "", "HTML file to parse")
	fs.StringVar(&opts.outputFile, "outputFile", "", "JSON file output")

	fs.Parse(args)

	if opts.inputFile == "" {
		return opts, errors.New("must provide an input file")
	}
	if opts.outputFile == "" {
		opts.outputFile = opts.inputFile + ".json"
	}
	return opts, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Start")

	opts, err := getOptions(os.Args[1:])
	if err != nil {
		log.Fatalf("Flag parsing error: %s", err.Error())
	}

	bytes, err := os.ReadFile(opts.inputFile)
	if err != nil {
		log.Fatalf("Can't open %s", opts.inputFile)
	}

	e := linklib.NewExtractor()

	ll := e.Extract(bytes)

	json_encoding, err := ll.ToJSON()
	if err != nil {
		log.Fatal("Can't encode to JSON")
	}

	log.Printf("Writing to %s", opts.outputFile)
	log.Printf("%s", json_encoding)
	os.WriteFile(opts.outputFile, json_encoding, 0666)
}
