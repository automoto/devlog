package pkg

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
	"time"
)

// TextContent interface for text content generation
type TextContent interface {
	GenerateMarkdown() string
}

// TemplateReader interface for reading templates
type TemplateReader interface {
	ReadTemplate() (*bytes.Buffer, error)
}

// Content struct which has attributes we need for generating the document
type Content struct {
	FormattedCurrentTime string
	TemplatePath         string
	DocumentType         string
}

func getTemplate(docType string) (string, error) {
	if docType == "note" {
		return defaultTemplate, nil
	} else if docType == "log" {
		return logTemplate, nil
	} else if docType == "todo" {
		return tdTemplate, nil
	}
	return "", errors.New("template not found for document type")
}

// ReadTemplate get a template based on the configured options and read it
func (c Content) ReadTemplate() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	documentTemplate, err := getTemplate(c.DocumentType)
	handleError(err)
	tpl := template.Must(template.New("devlog").Parse(documentTemplate))

	if len(c.TemplatePath) > 1 {
		tpl = template.Must(template.ParseFiles(c.TemplatePath))
	}
	err = tpl.Execute(buf, c)
	return buf, err
}

// GenerateMarkdown takes the output from a template and outputs it into a string
func (c Content) GenerateMarkdown() string {
	buff, err := c.ReadTemplate()
	if err != nil {
		handleError(err)
	}
	return buff.String()
}

func getTemplatePath(tmpl string, docType string) string {
	if len(tmpl) >= 1 {
		return tmpl
	}
	path := ""
	if docType == "note" {
		path = os.Getenv("DEVLOG_NOTE_TEMPLATE")
	} else if docType == "todo" {
		path = os.Getenv("DEVLOG_TODO_TEMPLATE")
	} else if docType == "log" {
		path = os.Getenv("DEVLOG_LOG_TEMPLATE")
	}
	return path
}

func getTrimmedOutput(output string) string {
	return strings.Trim(output, " ")
}

func checkStdOut(output string) bool {
	trimmedOutput := getTrimmedOutput(output)
	lowerOutPut := strings.ToLower(trimmedOutput)
	contains := strings.Contains(lowerOutPut, "stdout")
	return contains
}

func getOutputPath(outputFilePath string) string {
	if len(outputFilePath) >= 1 {
		return outputFilePath
	}
	envVarPath := os.Getenv("DEVLOG_DIR")
	if len(envVarPath) >= 1 {
		return envVarPath
	}
	return "."
}

func getFullOutputPath(outputFilePath string, docType string) string {
	return fmt.Sprintf("%s/%s", getOutputPath(outputFilePath), generateFileName(docType))
}

func generateFileName(docType string) string {
	now := time.Now()
	return fmt.Sprintf("devlog_%s_%s_%d-%d-%d.md", docType, now.Format("01_02_2006"), now.Hour(),
		now.Minute(), now.Second())
}

func saveFile(outputMd string, file io.Writer, outputFilePath string, docType string) {
	_, err := fmt.Fprint(file, outputMd)
	handleError(err)
	fmt.Println("Successfully saved dev log to path: ")
	fmt.Printf("%s\n", getFullOutputPath(outputFilePath, docType))
}
