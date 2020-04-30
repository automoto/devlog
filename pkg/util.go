package pkg

import (
	"fmt"
	"log"
	"os"
	"time"
)

func getCurrentDayAndTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func getCurrentDay() string {
	return time.Now().Format("2006-01-02")
}

func archive() {
	fmt.Println("not implemented yet")
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Devlog failed %v\n", err)
		log.Fatalln(err)
	}
}

// Start is the global executor that pulls in the configuration settings, generates the content and saves the file.
func Start(templatePath string, outputDirPath string) {
	var content contentConfig
	content.getContent(templatePath)
	output := generateMd(content.Questions, content.Other)
	if checkStdOut(outputDirPath) {
		fmt.Printf("%s", output)
	} else {
		file, err := os.Create(getFullOutputPath(outputDirPath))
		handleError(err)
		saveFile(output, file, outputDirPath)
	}
}
