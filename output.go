package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func generateMd(questions []string, otherSections []string) string {
	out := ""
	out += "### Dev Log\n"
	out += fmt.Sprintf("*created: %s*\n\n", getCurrentDayAndTime())
	for _, q := range questions {
		out += fmt.Sprintf("\n##### %s\n\n\n", q)
	}
	if len(otherSections) >= 1 {
		for _, q := range otherSections {
			out += fmt.Sprintf("\n##### %s\n\n\n", q)
		}
	}
	return out
}

func getLogContentPath() string {
	path := os.Getenv("DEVLOG_LOG_CONTENT")
	if len(path) > 1 {
		return path
	}
	return ""
}

type contentConfig struct {
	Questions []string `yaml:questions`
	Other     []string `yaml:other`
}

func getDefaultQuestions() string {
	return `
questions:
  - "How did your development session go? "
  - "Did you learn anything new? If so, what did you learn? "
  - "What could have gone better? "
  - "What went well? "
other:
  - "TODO"
  - "Notes"
`
}

func (c *contentConfig) getContent() *contentConfig {
	logPath := getLogContentPath()
	if len(logPath) >=1 {
		yamlFile, err := ioutil.ReadFile(logPath)
		handleError(err)
		err = yaml.Unmarshal(yamlFile, c)
		handleError(err)
		return c
	}

	err := yaml.UnmarshalStrict([]byte(getDefaultQuestions()), c)
	handleError(err)
	return c
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
