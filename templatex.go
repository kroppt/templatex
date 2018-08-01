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
	var options struct {
		op    string
		fin   string
		fout  string
		conf  string
		human bool
		guide bool
	}
	flag.StringVar(&options.op, "op", "build", "operation to perform out of: build, compile")
	flag.StringVar(&options.fin, "in", "stdin", "input file")
	flag.StringVar(&options.fout, "out", "stdout", "output file")
	flag.StringVar(&options.conf, "config", "", "template configuration file")
	flag.BoolVar(&options.human, "h", false, "use human readable json")
	flag.BoolVar(&options.guide, "guided", false, "prompt the user for more input")
	flag.Parse()
	var fin *os.File
	var err error
	if options.fin == "stdin" {
		fin = os.Stdin
	} else {
		fin, err = os.Open(options.fin)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed opening input file \"%v\" for reading: %v\n", options.fin, err)
		os.Exit(1)
	}
	var fout *os.File

	if options.fout == "stdout" {
		fout = os.Stdout
	} else {
		fout, err = os.OpenFile(options.fout, os.O_CREATE|os.O_WRONLY, 0644)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed opening output file \"%v\" for reading/writing: %v\n", options.fout, err)
		os.Exit(1)
	}
	switch options.op {
	case "build":
		buf, err := latex.GetConfig(fin, options.human)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed reading template: %v\n", err)
			os.Exit(1)
		}
		_, err = fout.Write(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed writing to \"%v\": %v\n", options.fout, err)
			os.Exit(1)
		}
	case "compile":
		// if config file is default, print usage
		if options.conf == "" {
			flag.Usage()
			os.Exit(2)
		}
		conf, err := os.Open(options.conf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed reading config \"%v\": %v\n", options.conf, err)
			os.Exit(1)
		}
		err = latex.CompileTemplate(fin, fout, conf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed compiling compiling document: %v\n", err)
			os.Exit(1)
		}
	default:
		flag.Usage()
		// Exit code 2: The command line parameters could not be parsed.
		os.Exit(2)
	}
	os.Exit(0)
}
