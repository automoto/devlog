package main

import (
	"flag"
	"fmt"
	"github.com/automoto/devlog/pkg"
)

var devlogVersion string

func main() {
	pathDescription := "Path to the directory where devlog will save notes to. Only pass in a path to a directory, the file names will be auto-generated."
	typeDescription := "Document type you wish to generate. Valid options are: todo, note, log"
	outputFlagPtr := flag.String("p", "", pathDescription)
	flag.String("path", "", pathDescription)
	typeFlagPtr := flag.String("t", "note", typeDescription)
	flag.String("type", "note", typeDescription)
	templateFlagPtr := flag.String("template", "", "File path to the .gohtml file for customizing your devlog notes.")
	versionFlagPtr := flag.Bool("v", false, "Return the current version of devlog")

	flag.Parse()
	if *versionFlagPtr {
		fmt.Println(devlogVersion)
		return
	}
	pkg.Start(*templateFlagPtr, *outputFlagPtr, *typeFlagPtr)
}
