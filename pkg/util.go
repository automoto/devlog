package pkg

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Time interface {
	GetCurrentDayAndTime() time.Time
}

type CurrentTime struct {}

func (c CurrentTime) GetCurrentDayAndTime() time.Time {
	return time.Now()
}

func archive() {
	fmt.Println("not implemented yet")
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}

// Start is the global executor that pulls in the configuration settings, generates the content and saves the file.
func Start(templatePath string, outputDirPath string) {
	ct := CurrentTime{}
	c := Content{
		 FormattedCurrentTime: ct.GetCurrentDayAndTime().Format("2006-01-02 15:04:05"),
		 TemplatePath:         templatePath,
	 }
	output := c.GenerateMarkdown()
	if checkStdOut(outputDirPath) {
		fmt.Printf("%s", output)
	} else {
		file, err := os.Create(getFullOutputPath(outputDirPath))
		handleError(err)
		saveFile(output, file, outputDirPath)
	}
}
