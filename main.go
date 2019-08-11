package main

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
)

func main() {
	prompt := promptui.Select{
		Label: "How did you feel",
		Items: []string{
			"Average: Kind of distracted, still got some things done, slightly tired but still have some energy",
			"Focused: Productive, not very distracted and energetic",
			"Distracted: Very Distracted, Didn't Get Alot Of Coding Done",
			"Tired: Too Mentally or Physically Exhausted To Focus",
			"Other: "},
	}

	result := runSelectPrompt(prompt)

	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt2 := promptui.Prompt{
		Label: "How long did you design and code for in minutes",
		Validate: validate,
	}
	result2 := runInputPrompt(prompt2)
	fmt.Printf("You choose %q\n", result)
	fmt.Printf("You choose %q\n", result2)
}

func runInputPrompt(p promptui.Prompt) string{
	result, err2 := p.Run()
	handleError(err2)
	return result
}

func runSelectPrompt(p promptui.Select) string{
	_, result, err := p.Run()
	handleError(err)
	return result
}

func runNumSelectPrompt(p promptui.Select) int{
	_, result, err := p.Run()
	handleError(err)
	parseResult, err := strconv.Atoi(result)
	handleError(err)
	return parseResult
}
