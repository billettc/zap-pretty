package main

import (
	"bufio"
	"flag"
	"os"

	zappretty "github.com/maoueh/zap-pretty"
)

// Provided via ldflags by goreleaser automatically

var showAllFlag = flag.Bool("all", false, "Show ")
var versionFlag = flag.Bool("version", false, "Prints version information and exit")

var showAll = false

func main() {
	flag.Parse()

	if *versionFlag {
		zappretty.PrintVersion()
		os.Exit(0)
	}

	go zappretty.NewSignaler().ForwardAllSignalsToProcessGroup()

	// FIXME: How could we make it more resilient to we simply drop the line instead? Would that mean our own "scanner"?
	// New scanner with a maximum of 250MiB per line, pass that, we panic.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(nil, 250*1024*1024)

	processor := zappretty.NewProcessor(scanner, os.Stdout, *showAllFlag)

	processor.Process()
}
