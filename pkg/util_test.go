package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetCurrentDayAndTime(t *testing.T) {
	ct := CurrentTime{}
	t.Run("generates a date and time", func(t *testing.T) {
		got := ct.GetCurrentDayAndTime()
		assert.IsType(t, time.Time{}, got)
	})
}
