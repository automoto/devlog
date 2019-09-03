package main

import (
	"fmt"
	"io/ioutil"
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

func (q *contentConfig) getContent() *contentConfig {
	yamlFile, err := ioutil.ReadFile("questions.yaml")
	handleError(err)
	err = yaml.UnmarshalStrict(yamlFile, q)
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
