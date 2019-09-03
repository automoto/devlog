package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func generateMd(questions []string, otherSections []string) string {
	out := ""
	out += "### Dev Log\n"
	out += fmt.Sprintf("*created: %s*", getCurrentDayAndTime())
	for _, q := range questions {
		out += fmt.Sprintf("\n##### %s\n", q)
	}
	for _, q := range otherSections {
		out += fmt.Sprintf("\n##### %s\n", q)
	}
	return out
}

func getOutputPath() string {
	path := os.Getenv("DEVLOG_DIR")
	if len(path) > 1 {
		return path
	}
	return "."
}

func getFullOutputPath() string {
	return fmt.Sprintf("%s/%s", getOutputPath(), generateFileName())
}

func generateFileName() string {
	now := time.Now()
	return fmt.Sprintf("devlog_%d_%d_%d_%d_%d_%d.md", now.Day(), now.Month(), now.Year(), now.Hour(),
		now.Minute(), now.Second())
}

func saveFile(outputMd string, file io.Writer) {
	_, err := fmt.Fprint(file, outputMd)
	handleError(err)
	log.Printf("Successfully saved dev log to directory: %s", getOutputPath())
}
