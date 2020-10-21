package main

import (
	"flag"
	"fmt"
	"github.com/automoto/devlog/pkg"
)

var devlogVersion string

func main() {
	outputFlagPtr := flag.String("p", "", "Path to the directory where devlog will save notes to. Only pass in a path to a directory, the file names will be auto-generated.")
	templateFlagPtr := flag.String("t", "", "File path to the .gohtml file for customizing your devlog notes.")
	versionFlagPtr := flag.Bool("v", false, "Return the current version of devlog")
	flag.Parse()
	if *versionFlagPtr {
		fmt.Println(devlogVersion)
		return
	}
	pkg.Start(*templateFlagPtr, *outputFlagPtr)
}
