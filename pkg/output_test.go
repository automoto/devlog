package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateMd(t *testing.T) {
	testQuestions := []string{"question 1", "question 2"}
	otherSections := []string{"section 1", "section 2"}
	t.Run("Markdown gets generated correctly", func(t *testing.T) {
		got := generateMd(testQuestions, otherSections)
		assert.NotEmpty(t, got)
		assert.Contains(t, got, "##### section 1")
		assert.Contains(t, got, "##### section 2")
		assert.Contains(t, got, "##### question 1")
		assert.Contains(t, got, "##### question 2")
	})
}
