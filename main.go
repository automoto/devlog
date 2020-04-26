package main

import (
	"flag"
)

func main() {
	outputFlagPtr := flag.String("f", "", "path to the directory where devlog will save notes to")
	templateFlagPtr := flag.String("c", "", "config for your devlog notes")
	flag.Parse()
	start(*templateFlagPtr, *outputFlagPtr)
}
