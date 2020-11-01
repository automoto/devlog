package pkg

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentDayAndTime(t *testing.T) {
	ct := CurrentTime{}
	t.Run("generates a date and time", func(t *testing.T) {
		got := ct.GetCurrentDayAndTime()
		assert.IsType(t, time.Time{}, got)
	})
}

func TestContains(t *testing.T) {
	type args struct {
		slice []string
		val   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns true when a value found",
			args: args{
				slice: []string{"a", "b", "c"},
				val: "b",
			},
			want: true,
		},
		{
			name: "returns false when a value not found",
			args: args{
				slice: []string{"a", "b", "c"},
				val: "z",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.slice, tt.args.val); got != tt.want {
				t.Errorf("Contains(%v, %v) = %v, want %v", tt.args.slice, tt.args.val, got, tt.want)
			}
		})
	}
}

func Test_validDocTypes(t *testing.T) {
	t.Run("expected number of valid document types is returned", func(t *testing.T) {
		got := validDocTypes()
		assert.Len(t, got, 3)
	})
}

func Test_isDocTypeValid(t *testing.T) {
	t.Run("returns correct values when document type not found", func(t *testing.T) {
		got, err := isDocTypeValid("fgfdhglkdsf")
		assert.Error(t, err)
		assert.False(t, got)
	})
	t.Run("returns correct values when document type is found", func(t *testing.T) {
		got, err := isDocTypeValid("note")
		assert.NoError(t, err)
		assert.True(t, got)
	})
}

func Test_cleanInput(t *testing.T) {
	t.Run("cleans input", func(t *testing.T) {
		testInput := " hello Bill "
		expectedOutput := "hello bill"
		got := cleanInput(testInput)
		assert.Equal(t, expectedOutput, got)
	})
}
