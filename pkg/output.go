package pkg

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
	"time"
)

// interface for content

type TextContent interface {
	GenerateMarkdown() string
}

type TemplateReader interface {
	ReadTemplate() (*bytes.Buffer, error)
}

type Content struct {
	FormattedCurrentTime string
	TemplatePath         string
}

func (c Content) ReadTemplate() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	tpl := template.Must(template.New("devlog").Parse(defaultTemplate))

	if len(c.TemplatePath) > 1 {
		tpl = template.Must(template.ParseFiles(c.TemplatePath))
	}
	err := tpl.Execute(buf, c)
	return buf, err
}

func (c Content) GenerateMarkdown() string {
	buff, err := c.ReadTemplate()
	if err != nil {
		handleError(err)
	}
	return buff.String()
}

func (c Content) GetTemplatePath() string {
	if len(c.TemplatePath) >= 1 {
		return c.TemplatePath
	}
	path := os.Getenv("DEVLOG_LOG_CONTENT")
	if len(path) >= 1 {
		return path
	}
	return ""
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

func getFullOutputPath(outputFilePath string) string {
	return fmt.Sprintf("%s/%s", getOutputPath(outputFilePath), generateFileName())
}

func generateFileName() string {
	now := time.Now()
	return fmt.Sprintf("devlog_%s_%d-%d-%d.md", now.Format("01_02_2006"), now.Hour(),
		now.Minute(), now.Second())
}

func saveFile(outputMd string, file io.Writer, outputFilePath string) {
	_, err := fmt.Fprint(file, outputMd)
	handleError(err)
	fmt.Println("Successfully saved dev log to path: ")
	fmt.Printf("%s\n", getFullOutputPath(outputFilePath))
}

//TODO: add back .yaml based config for default options
//type DevLogConfig struct {
//	Questions []string `yaml:"questions"`
//	Other     []string `yaml:"other_section"`
//}

//func (c *DevLogConfig) getConfig(configFilePath string) *DevLogConfig {
//	logPath := getLogContentPath(templateFilePath)
//	if len(logPath) >=1 {
//		yamlFile, err := ioutil.ReadFile(logPath)
//		handleError(err)
//		err = yaml.Unmarshal(yamlFile, c)
//		handleError(err)
//		return c
//	}
//
//	err := yaml.UnmarshalStrict([]byte(getDefaultQuestions()), c)
//	handleError(err)
//	return c
//}
