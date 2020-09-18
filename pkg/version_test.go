package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVersion(t *testing.T) {
	t.Run("Gets the version", func(t *testing.T) {
		assert.Equal(t, "v0.0.8", getVersion())
	})
}
