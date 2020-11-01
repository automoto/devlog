package pkg

import (
	"bytes"
	"errors"
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
			c := Content{
				DocumentType: tt.args.docType,
			}
			got, err := c.GetTemplate()
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

func fakeEnvVarFetcher(envVar string) string {
	if envVar == "DEVLOG_NOTE_TEMPLATE" {
		return "/home/note"
	} else if envVar == "DEVLOG_TODO_TEMPLATE" {
		return "/home/todo"
	} else if envVar == "DEVLOG_LOG_TEMPLATE" {
		return "/home/log"
	}
	return ""
}

func Test_GetTemplatePath(t *testing.T) {
	osGetEnv = fakeEnvVarFetcher
	type args struct {
		tmpl    string
		docType string
		envVarKeySet string
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
				envVarKeySet: "fake",
			},
			want: "",
		},
		{
			name: "env var path for note is returned when no template path in input present and docType is note",
			args: args{
				tmpl:    "",
				docType: "note",
				envVarKeySet: "DEVLOG_NOTE_TEMPLATE",
			},
			want: "/home/note",
		},
		{
			name: "env var path for todo is returned when no template path in input present and docType is todo",
			args: args{
				tmpl:    "",
				docType: "todo",
				envVarKeySet: "DEVLOG_TODO_TEMPLATE",
			},
			want: "/home/todo",
		},
		{
			name: "env var path for log is returned when no template path in input present and docType is log",
			args: args{
				tmpl:    "",
				docType: "log",
				envVarKeySet: "DEVLOG_LOG_TEMPLATE",
			},
			want: "/home/log",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Content{
				FormattedCurrentTime: testTime,
				DocumentType:         tt.args.docType,
			}
			got := c.GetTemplatePath(tt.args.tmpl)
			assert.Equal(t, tt.want, got)
		})
	}
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

func fakeGetEnvSet(envVar string) string{
	return "/home/test"
}

func fakeGetEnvNotSet(envVar string) string {
	return ""
}

func Test_GetOutputPath(t *testing.T) {
	t.Run("output path is returned when path is set via input", func(t *testing.T) {
		df := DevlogFile{
			OutputDirPath: "/home/kanye",
		}
		got := df.GetOutputPath()
		assert.Equal(t, "/home/kanye", got)
	})
	t.Run("environment variable path returned when no input and environment variable is set", func(t *testing.T) {
		df := DevlogFile{}
		osGetEnv = fakeGetEnvSet
		got := df.GetOutputPath()
		assert.Equal(t, "/home/test", got)
	})
	t.Run("environment variable path returned when no input and environment variable is set", func(t *testing.T) {
		df := DevlogFile{}
		osGetEnv = fakeGetEnvNotSet
		got := df.GetOutputPath()
		assert.Equal(t, ".", got)
	})
}

func Test_GenerateFileName(t *testing.T) {
	t.Run("filename gets returned", func(t *testing.T) {
		df := DevlogFile{}
		got := df.GenerateFileName("note")
		assert.Contains(t, got, "devlog_note_")
	})
}

type fakeOS struct {
	name string
	err error
}
func (f *fakeOS) fakeCreate(name string) ( *os.File, error){
	f.name = name
	testFile := os.File{}
	return &testFile, f.err
}

func TestDevlogFile_CreateFile(t *testing.T) {
	f := &fakeOS{
		name: "testFile",
		err:  nil,
	}
	osCreate = f.fakeCreate
	df := DevlogFile{
		OutputFilePath: "/home/test",
	}
	t.Run("creates file to be saved without error", func(t *testing.T) {
		_, err := df.CreateFile()
		assert.NoError(t, err)
	})
	t.Run("returns an error when file creation fails", func(t *testing.T) {
		f := &fakeOS{
			name: "error",
			err:  errors.New("failed to create file"),
		}
		osCreate = f.fakeCreate
		_, err := df.CreateFile()
		assert.Error(t, err)
	})
}

func TestDevlogFile_SaveFile(t *testing.T) {
	t.Run("saves output", func(t *testing.T) {
		df := DevlogFile{
			OutputFilePath: "/home/fly",
		}
		buffer := bytes.Buffer{}
		err := df.SaveFile("#Note +1", &buffer, "note")
		assert.NoError(t, err)
	})
}