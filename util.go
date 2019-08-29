package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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

func promptMoodChoices(label string, choices []string) promptui.Select{
	return promptui.Select{
		Label: label,
		Items: choices,
	}
}

func promptInputNoValidate(label string) promptui.Prompt {
	return promptui.Prompt{
		Label: label,
	}
}

func getAnswers(selectChoices []string, questions []string) map[string]string {
	questionAnswerPairs := make(map[string]string)
	labelSelect := "How did you feel? "
	questionAnswerPairs[labelSelect] = getSelectInput(promptMoodChoices(labelSelect, selectChoices))

	for _, q := range questions {
		questionAnswerPairs[q] = getInput(promptInputNoValidate(q))
	}
	return questionAnswerPairs
}

func generateMd(selectChoices []string, questions []string) string{
	out := ""
	qa := getAnswers(selectChoices, questions)
	for q, a := range qa {
		out += fmt.Sprintf("\n#### %s\n%s\n", q, a)
	}
	return out
}

type listOfQuestions struct {
	Status []string `yaml:status`
	Questions []string `yaml:questions`
}

func (q *listOfQuestions) getQuestions() *listOfQuestions{
	yamlFile, err := ioutil.ReadFile("questions.yaml")
	handleError(err)
	err = yaml.Unmarshal(yamlFile, q)
	handleError(err)
	fmt.Println(q)
	return q
}

func startPrompt() {
	var questionsToPrompt listOfQuestions
	questionsToPrompt.getQuestions()
	output := generateMd(questionsToPrompt.Status, questionsToPrompt.Questions)
	fmt.Println(output)
}

func getSelectInput(p promptui.Select) string{
	_, result, err := p.Run()
	handleError(err)
	return result
}

func getInput(p promptui.Prompt) string{
	result, err2 := p.Run()
	handleError(err2)
	return result
}
