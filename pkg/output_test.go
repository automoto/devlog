package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestReadTemplate(t *testing.T) {
	c := Content{
		FormattedCurrentTime: time.Now().Format("2006-01-02 15:04:05"),
		TemplatePath:         "",
	}
	t.Run("Template gets generated with the default template", func(t *testing.T) {
		got, err := c.ReadTemplate()
		sg := got.String()
		assert.NoError(t, err)
		assert.Contains(t, sg, "### Development Log")
		assert.Contains(t, sg, "##### What could have gone better?")
		assert.Contains(t, sg, c.FormattedCurrentTime)
	})
	//TODO: More of an integration test here, consider mocking the file system and moving this out to a separate integration test
	t.Run("Template gets generated with the default template", func(t *testing.T) {
		c.TemplatePath = "test_template.gohtml"
		got, err := c.ReadTemplate()
		sg := got.String()
		assert.NoError(t, err)
		assert.Contains(t, sg, "### Test Log")
		assert.Contains(t, sg, c.FormattedCurrentTime)
	})
}
