package main

import (
	"flag"
)

func main() {
	outputFlagPtr := flag.String("p", "", "path to the directory where devlog will save notes to")
	templateFlagPtr := flag.String("t", "", "template for your devlog notes")
	flag.Parse()
	start(*templateFlagPtr, *outputFlagPtr)
}
