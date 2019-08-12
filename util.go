package main

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
)

func getCurrentDay() {

}

func generateOutput() {

}

func archive() {
	fmt.Println("not implemented yet")
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err)
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

func generateMd(questions []string, time string, learned string, activities string, toImprove string, wentWell string) string {
	return ""
}

func startPrompt() {
	output := generateMd(
		getSelectInput(promptMoodChoices()),
		getInput(promptInput(validateNumFunc(), "How long, in minutes, did you design and code for? ")),
		getInput(promptInputNoValidate("What did you learn? ")),
		getInput(promptInputNoValidate("What did you do? ")),
		getInput(promptInputNoValidate("What could have gone better? ")),
		getInput(promptInputNoValidate("What went well? ")),
	)
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
