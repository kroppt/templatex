package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kroppt/templatex/latex"
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
		human bool
		guide bool
	}
	flag.StringVar(&config.op, "op", "build", "operation to perform out of: build, compile")
	flag.StringVar(&config.fin, "in", "stdin", "input file")
	flag.StringVar(&config.fout, "out", "stdout", "output file")
	flag.BoolVar(&config.human, "h", false, "human readable json")
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
	var fin *os.File
	var err error
	if config.fin == "stdin" {
		fin = os.Stdin
	} else {
		fin, err = os.Open(config.fin)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "input file could not be opened for reading")
		os.Exit(1)
	}
	var fout *os.File
	if config.fout == "stdout" {
		fout = os.Stdout
	} else {
		fout, err = os.OpenFile(config.fout, os.O_CREATE|os.O_WRONLY, 0644)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "output file could not be opened for reading and writing")
		os.Exit(1)
	}
	buf, err := latex.GetConfig(fin, config.human)
	fout.Write(buf)
	os.Exit(0)
}
