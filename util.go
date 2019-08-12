package main

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strconv"
)

func getCurrentDay() {

}

func archive() {
	fmt.Println("not implemented yet")
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		log.Fatalln(err)
	}
}

func promptMoodChoices() promptui.Select{
	return promptui.Select{
		Label: "How did the session go",
		Items: []string{
			"Average: Kind of distracted, still got some things done, slightly tired but still have some energy",
			"Focused: Productive, not very distracted and energetic",
			"Distracted: Very Distracted, Didn't Get Alot Of Coding Done",
			"Tired: Too Mentally or Physically Exhausted To Focus",
			"Other: "},
	}
}

func promptInput(validateNum promptui.ValidateFunc, label string) promptui.Prompt {
	return promptui.Prompt{
		Label: label,
		Validate: validateNum,
	}
}

func promptInputNoValidate(label string) promptui.Prompt {
	return promptui.Prompt{
		Label: label,
	}
}

func validateNumFunc() promptui.ValidateFunc{
	return func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("invalid number")
		}
		return nil
	}
}

func getAnswers(questions []string) map[string]string {
	questionAnswerPairs := make(map[string]string)
	for _, q := range questions {
		questionAnswerPairs[q] = getInput(promptInputNoValidate(q))
	}
	return questionAnswerPairs
}

func generateMd(questions []string) string{
	out := ""
	qa := getAnswers(questions)
	for q, a := range qa {
		out += fmt.Sprintf("\n#### %s\n%s\n", q, a)
	}
	return out
}

type listOfQuestions struct {
	Questions []string `yaml:questions`
}

func (q *listOfQuestions) getQuestions() *listOfQuestions{
	yamlFile, err := ioutil.ReadFile("questions.yaml")
	handleError(err)
	err = yaml.Unmarshal(yamlFile, q)
	handleError(err)
	return q
}

func startPrompt() {
	var questionsToPrompt listOfQuestions
	questionsToPrompt.getQuestions()
	output := generateMd(questionsToPrompt.Questions)
	fmt.Println(output)
}

func getInput(p promptui.Prompt) string{
	result, err2 := p.Run()
	handleError(err2)
	return result
}

func getSelectInput(p promptui.Select) string{
	_, result, err := p.Run()
	handleError(err)
	return result
}
