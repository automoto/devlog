package pkg

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// test data
var testTime = time.Now().Format("2006-01-02 15:04:05")
var expectedLogOutput = fmt.Sprintf(`### Development Log
*created: %s*


##### How did your development session go?



##### Did you learn anything new? If so, what did you learn?



##### What could have gone better?



##### What went well?


---

##### Notes

`, testTime)
var expectedTdOutput = fmt.Sprintf(`### TODO
*created: %s*

- [ ]
- [ ]

`, testTime)
var expectedDefaultOutput = fmt.Sprintf(`### Note
*created: %s*

`, testTime)

func TestReadTemplate(t *testing.T) {
	c := Content{
		FormattedCurrentTime: testTime,
		TemplatePath:         "",
		DocumentType:         "note",
	}
	t.Run("Template gets generated correctly with the 'note' template", func(t *testing.T) {
		got, err := c.ReadTemplate()
		sg := got.String()
		assert.NoError(t, err)
		assert.Equal(t, expectedDefaultOutput, sg)
	})
	t.Run("Template gets generated correctly with the 'log' template", func(t *testing.T) {
		c.DocumentType = "log"
		got, err := c.ReadTemplate()
		sg := got.String()
		assert.NoError(t, err)
		assert.Equal(t, expectedLogOutput, sg)
	})
	t.Run("Template gets generated correctly with the 'todo' template", func(t *testing.T) {
		c.DocumentType = "todo"
		got, err := c.ReadTemplate()
		sg := got.String()
		assert.NoError(t, err)
		assert.Equal(t, expectedTdOutput, sg)
	})
	//TODO: More of an integration test here, consider mocking the file system and moving this out to a separate integration test
	t.Run("Template gets generated with a custom template", func(t *testing.T) {
		c.TemplatePath = "test_template.gohtml"
		got, err := c.ReadTemplate()
		sg := got.String()
		assert.NoError(t, err)
		assert.Contains(t, sg, "### Test Log")
		assert.Contains(t, sg, c.FormattedCurrentTime)
	})
}

func Test_getTemplate(t *testing.T) {
	type args struct {
		docType string
	}
	tests := []struct {
		name     string
		args     args
		contains string
		wantErr  bool
	}{
		{
			name: "todo template is found",
			args: args{
				docType: "todo",
			},
			contains: "### TODO",
			wantErr: false,
		},
		{
			name: "note template is found",
			args: args{
				docType: "note",
			},
			contains: "### Note",
			wantErr: false,
		},
		{
			name: "log template is found",
			args: args{
				docType: "log",
			},
			contains: "### Development Log",
			wantErr: false,
		},
		{
			name: "error for non-existent template",
			args: args{
				docType: "foobar",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTemplate(tt.args.docType)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTemplate(%v) error = %v, wantErr %v", tt.args.docType, err, tt.wantErr)
				return
			}
			assert.Contains(t, got, tt.contains)
		})
	}
}

func TestContent_ReadTemplate(t *testing.T) {
	c := Content{
		FormattedCurrentTime: testTime,
		TemplatePath:         "",
		DocumentType:         "todo",
	}
		t.Run("Read template returns without an error with no template", func(t *testing.T) {
			got, err := c.ReadTemplate()
			assert.NoError(t, err)
			assert.NotEmpty(t, got.String())
		})
	// TODO: More of an integration test, consider mocking or moving to integration test suite
	t.Run("Read template returns without an error with a template", func(t *testing.T) {
		c.TemplatePath = "test_template.gohtml"
		got, err := c.ReadTemplate()
		assert.NoError(t, err)
		assert.NotEmpty(t, got.String())
	})
}

func TestContent_GenerateMarkdown(t *testing.T) {
	t.Run("Returns a string", func(t *testing.T){
		c := Content{
			FormattedCurrentTime: testTime,
			TemplatePath:         "",
			DocumentType:         "todo",
		}
		got := c.GenerateMarkdown()
		assert.NotEmpty(t, got)
	})
}

func Test_getTemplatePath(t *testing.T) {
	type args struct {
		tmpl    string
		docType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTemplatePath(tt.args.tmpl, tt.args.docType); got != tt.want {
				t.Errorf("getTemplatePath(%v, %v) = %v, want %v", tt.args.tmpl, tt.args.docType, got, tt.want)
			}
		})
	}
}

func Test_getTrimmedOutput(t *testing.T) {
	type args struct {
		output string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTrimmedOutput(tt.args.output); got != tt.want {
				t.Errorf("getTrimmedOutput(%v) = %v, want %v", tt.args.output, got, tt.want)
			}
		})
	}
}

func Test_checkStdOut(t *testing.T) {
	type args struct {
		output string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStdOut(tt.args.output); got != tt.want {
				t.Errorf("checkStdOut(%v) = %v, want %v", tt.args.output, got, tt.want)
			}
		})
	}
}

func Test_getOutputPath(t *testing.T) {
	type args struct {
		outputFilePath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOutputPath(tt.args.outputFilePath); got != tt.want {
				t.Errorf("getOutputPath(%v) = %v, want %v", tt.args.outputFilePath, got, tt.want)
			}
		})
	}
}

func Test_getFullOutputPath(t *testing.T) {
	type args struct {
		outputFilePath string
		docType        string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFullOutputPath(tt.args.outputFilePath, tt.args.docType); got != tt.want {
				t.Errorf("getFullOutputPath(%v, %v) = %v, want %v", tt.args.outputFilePath, tt.args.docType, got, tt.want)
			}
		})
	}
}

func Test_generateFileName(t *testing.T) {
	type args struct {
		docType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateFileName(tt.args.docType); got != tt.want {
				t.Errorf("generateFileName(%v) = %v, want %v", tt.args.docType, got, tt.want)
			}
		})
	}
}

//func Test_saveFile(t *testing.T) {
//	type args struct {
//		outputMd       string
//		outputFilePath string
//		docType        string
//	}
//	tests := []struct {
//		name     string
//		args     args
//		wantFile string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			file := &bytes.Buffer{}
//			saveFile(tt.args.outputMd, file, tt.args.outputFilePath, tt.args.docType)
//			if gotFile := file.String(); gotFile != tt.wantFile {
//				t.Errorf("saveFile(%v, %v, %v, %v) = %v, want %v", tt.args.outputMd, tt.args.file, tt.args.outputFilePath, tt.args.docType, gotFile, tt.wantFile)
//			}
//		})
//	}
//}
