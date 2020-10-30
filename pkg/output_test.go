package pkg

import (
	"fmt"
	"os"
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
		envVars []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tmpl value is returned when its already set",
			args: args{
				tmpl: "/home/devlog",
				docType: "note",
			},
			want: "/home/devlog",
		},
		{
			name: "no path is returned when no template path present",
			args: args{
				tmpl:    "",
				docType: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTemplatePath(tt.args.tmpl, tt.args.docType); got != tt.want {
				t.Errorf("getTemplatePath(%v, %v) = %v, want %v", tt.args.tmpl, tt.args.docType, got, tt.want)
			}
		})
	}
	// TODO: More of an integration test, consider using mocks here.
	t.Run("set environment variable path is returned when present", func(t *testing.T) {
		os.Setenv("DEVLOG_NOTE_TEMPLATE", "/home/documents")
		got := getTemplatePath("", "note")
		assert.Equal(t, "/home/documents", got)
		os.Unsetenv("DEVLOG_NOTE_TEMPLATE")
	})
}

func Test_getTrimmedOutput(t *testing.T) {
	t.Run("output gets trimmed", func(t *testing.T) {
		got := getTrimmedOutput("billy ")
		assert.Equal(t, got, "billy")
	})
}

func Test_checkStdOut(t *testing.T) {
	t.Run("stdout is true when present", func(t *testing.T) {
		got := checkStdOut("stdout")
		assert.True(t, got)
	})
	t.Run("stdout is false when not present", func(t *testing.T) {
		got := checkStdOut("/home/")
		assert.False(t, got)
	})

}

func Test_getOutputPath(t *testing.T) {
	t.Run("output path is returned when present", func(t *testing.T) {
		got := getOutputPath("/home/kanye")
		assert.Equal(t, "/home/kanye", got)
	})
}

func Test_generateFileName(t *testing.T) {
	t.Run("filename gets returned", func(t *testing.T) {
		got := generateFileName("note")
		assert.Contains(t, got, "devlog_")
	})
}
