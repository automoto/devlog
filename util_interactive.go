package main

import "github.com/manifoldco/promptui"

func getAnswers(selectChoices []string, questions []string) map[string]string {
	questionAnswerPairs := make(map[string]string)
	labelSelect := "How did you feel? "
	questionAnswerPairs[labelSelect] = getSelectInput(promptMoodChoices(labelSelect, selectChoices))

	for _, q := range questions {
		questionAnswerPairs[q] = getInput(promptInputNoValidate(q))
	}
	return questionAnswerPairs
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
