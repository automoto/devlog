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

// Store func calls to os.Getenv in a variable to improve testability
var osGetEnv = os.Getenv

// TextContent interface for text content generation
type TextContent interface {
	GenerateMarkdown() string
}

// TemplateReader interface for reading templates
type TemplateReader interface {
	ReadTemplate() (*bytes.Buffer, error)
	GetTemplatePath() string
	GetTemplate() (string, error)
}

// GetTemplatePath gets the template path based on if an input is passed in or if an env var is set
func (c Content) GetTemplatePath(inputPath string) string {
	if len(inputPath) >= 1 {
		return inputPath
	}
	path := ""
	if c.DocumentType == "note" {
		path = osGetEnv("DEVLOG_NOTE_TEMPLATE")
	} else if c.DocumentType == "todo" {
		path = osGetEnv("DEVLOG_TODO_TEMPLATE")
	} else if c.DocumentType == "log" {
		path = osGetEnv("DEVLOG_LOG_TEMPLATE")
	}
	return path
}

// GetTemplate gets our document template based on the input document type
func (c Content) GetTemplate() (string, error) {
	if c.DocumentType == "note" {
		return defaultTemplate, nil
	} else if c.DocumentType == "log" {
		return logTemplate, nil
	} else if c.DocumentType == "todo" {
		return tdTemplate, nil
	}
	return "", errors.New("template not found for document type")
}

// Content struct which has attributes we need for generating the document
type Content struct {
	FormattedCurrentTime string
	TemplatePath         string
	DocumentType         string
	Tags                 string
}

// ReadTemplate get a template based on the configured options and read it
func (c Content) ReadTemplate() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	documentTemplate, err := c.GetTemplate()
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

func getTrimmedOutput(output string) string {
	return strings.Trim(output, " ")
}

func checkStdOut(output string) bool {
	trimmedOutput := getTrimmedOutput(output)
	lowerOutPut := strings.ToLower(trimmedOutput)
	contains := strings.Contains(lowerOutPut, "stdout")
	return contains
}

// DevlogFile holds metadata about files
type DevlogFile struct {
	OutputDirPath string
	OutputFilePath string
}

// FileOps is the interface for devlog file operations
type FileOps interface {
	GetOutputPath() string
	GetFullOutputPath(docType string) string
	GenerateFileName(docType string) string
	CreateFile(docType string) (*os.File, error)
	SaveFile(outputMd string, file io.Writer, docType string)
}

var osCreate = os.Create
// CreateFile creates a file for devlog to save
func (f DevlogFile) CreateFile() (*os.File, error){
	file, err := osCreate(f.OutputFilePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// SaveFile saves our markdown file to the specified directory
func (f DevlogFile) SaveFile(outputMd string, file io.Writer, docType string) error {
	_, err := fmt.Fprint(file, outputMd)
	if err != nil {
		return err
	}
	fmt.Println("Successfully saved dev log to path: ")
	fmt.Printf("%s\n", f.GetFullOutputPath(docType))
	return nil
}

// GetOutputPath gets a path and selects sensible defaults if one is not set
func (f DevlogFile) GetOutputPath() string {
	if len(f.OutputDirPath) >= 1 {
		return f.OutputDirPath
	}
	envVarPath := osGetEnv("DEVLOG_DIR")
	if len(envVarPath) >= 1 {
		return envVarPath
	}
	return "."
}

// GetFullOutputPath generates the full path with the specified directory and filename
func (f DevlogFile) GetFullOutputPath(docType string) string {
	return fmt.Sprintf("%s/%s", f.GetOutputPath(), f.GenerateFileName(docType))
}

// GenerateFileName generates the timestamped file prefixed with the document type
// TODO: refactor time.Now call here to use our time interface in util for testing
func (f DevlogFile) GenerateFileName(docType string) string {
	now := time.Now()
	return fmt.Sprintf("devlog_%s_%s_%d_%d_%d.md", docType, now.Format("01_02_2006"), now.Hour(),
		now.Minute(), now.Second())
}
