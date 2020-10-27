package pkg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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
