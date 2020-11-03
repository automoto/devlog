package pkg

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

// Time is our interface for any functions we need for handling time
type Time interface {
	GetCurrentDayAndTime() time.Time
}

// CurrentTime is our struct for holding the current time
type CurrentTime struct{}

// GetCurrentDayAndTime gets the current day and time and returns as a time.Time
func (c CurrentTime) GetCurrentDayAndTime() time.Time {
	return time.Now()
}

// Contains is a useful method for checking if a value exists in a slice
func Contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func validDocTypes() []string {
	validDocTypes := make([]string, 0)
	validDocTypes = append(validDocTypes, "note")
	validDocTypes = append(validDocTypes, "log")
	validDocTypes = append(validDocTypes, "todo")
	return validDocTypes
}

func isDocTypeValid(docTypeInput string) (bool, error) {
	if Contains(validDocTypes(), docTypeInput) {
		return true, nil
	}
	err := errors.New("document type not found. Valid document types are: note, todo, log")
	return false, err
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}

func cleanInput(inputString string) string {
	finalInput := ""
	finalInput = strings.ToLower(inputString)
	finalInput = strings.TrimSpace(finalInput)
	return finalInput
}

// Converts a string slice into a formatted string of tags
func ConvertTagsToString(tags []string) string {
	finalString := make([]string, 0)
	tagChar := "#"
	for _, value := range tags {
		trimmed := strings.TrimSpace(value)
		finalString = append(finalString, fmt.Sprintf("%s%s", tagChar, trimmed))
	}
	return strings.Join(finalString, " ")
}

// Parses the tags into a string slice
func ParseTags(inputTags string) []string {
	if len(inputTags) >= 1 {
		result := strings.Split(inputTags, ",")
		return result
	}
	return nil
}

// Start is the global executor that pulls in the configuration settings, generates the content and saves the file.
func Start(templatePath string, tags string, outputDirPath string, docType string) {
	docType = cleanInput(docType)
	_, err := isDocTypeValid(docType)
	handleError(err)
	ct := CurrentTime{}
	c := Content{
		FormattedCurrentTime: ct.GetCurrentDayAndTime().Format("2006-01-02 15:04:05"),
		TemplatePath:         templatePath,
		DocumentType:         docType,
		Tags:                 "",
	}
	parsedTags := ParseTags(tags)
	if len(parsedTags) >= 1 {
		c.Tags = ConvertTagsToString(parsedTags)
	}
	output := c.GenerateMarkdown()

	if checkStdOut(outputDirPath) {
		fmt.Printf("%s", output)
	} else {
		df := DevlogFile{
			OutputDirPath:  outputDirPath,
		}
		df.OutputFilePath = df.GetFullOutputPath(docType)
		file, err := df.CreateFile()
		handleError(err)
		err = df.SaveFile(output, file, docType)
		handleError(err)
	}
}
