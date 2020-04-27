package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCurrentDayAndTime (t *testing.T) {
	t.Run("generates a date and time", func(t *testing.T) {
		got := getCurrentDayAndTime()
		dateLen := len(got)
		assert.NotEmpty(t, got)
		assert.True(t, dateLen > 1)
	})
}
