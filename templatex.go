package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: templatex [flags]")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	var config struct {
		op    string
		fin   string
		fout  string
		guide bool
	}
	flag.StringVar(&config.op, "op", "build", "operation to perform out of: build, compile")
	flag.StringVar(&config.fin, "in", "stdin", "input file")
	flag.StringVar(&config.fout, "out", "stdout", "output file")
	flag.BoolVar(&config.guide, "guided", false, "prompt the user for more input")
	flag.Parse()
	switch config.op {
	case "build":
	case "compile":
	default:
		flag.Usage()
		// Exit code 2: The command line parameters could not be parsed.
		os.Exit(2)
	}
	_, err := os.Open(config.fin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "input file could not be opened for reading")
		os.Exit(1)
	}
	_, err = os.OpenFile(config.fout, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "output file could not be opened for reading and writing")
		os.Exit(1)
	}
}
