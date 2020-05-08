package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestReadTemplate(t *testing.T) {
	c := Content{
		CurrentTime:  time.Now(),
		TemplatePath: "",
	}
	t.Run("Template gets generated correctly with the default template", func(t *testing.T) {
		got, err := c.ReadTemplate()
		sg := got.String()
		assert.NoError(t, err)
		assert.Contains(t, sg, "### Development Log")
		assert.Contains(t, sg, "##### What could have gone better?")
	})
	t.Run("Template gets generated correctly with the default template", func(t *testing.T) {
		c.TemplatePath = "test_template.gohtml"
		got, err := c.ReadTemplate()
		sg := got.String()
		assert.NoError(t, err)
		assert.Contains(t, sg, "### Test Log")
	})
}
