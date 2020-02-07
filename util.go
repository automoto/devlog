package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
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

func (q *contentConfig) getContent() *contentConfig {
	err := yaml.UnmarshalStrict([]byte(getDefaultQuestions()), q)
	handleError(err)
	return q
}

func start() {
	var questions contentConfig
	questions.getContent()
	output := generateMd(questions.Questions, questions.Other)
	file, err := os.Create(getFullOutputPath())
	handleError(err)
	saveFile(output, file)
}
