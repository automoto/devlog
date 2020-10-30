package main

import (
	"flag"
	"fmt"
	"github.com/automoto/devlog/pkg"
)

var devlogVersion string

func main() {
	var pathFlagInput string
	var typeFlagInput string
	flag.StringVar(&pathFlagInput, "path", "", "Path to the directory where devlog will save notes to. Only pass in a path to a directory, the file names will be auto-generated.")
	flag.StringVar(&pathFlagInput, "p", "", "Shortcut for path.")
	flag.StringVar(&typeFlagInput, "type", "note", "Document type you wish to generate. Valid options are: todo, note, log ")
	flag.StringVar(&typeFlagInput, "t", "note", "Shortcut for document type.")
	templateFlagPtr := flag.String("template", "", "File path to the .gohtml file for customizing your devlog notes.")
	versionFlagPtr := flag.Bool("v", false, "Return the current version of devlog.")

	flag.Parse()
	if *versionFlagPtr {
		fmt.Println(devlogVersion)
		return
	}
	pkg.Start(*templateFlagPtr, pathFlagInput, typeFlagInput)
}
