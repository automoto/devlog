package pkg

import (
	"reflect"
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

func TestCurrentTime_GetCurrentDayAndTime(t *testing.T) {
	tests := []struct {
		name string
		c    CurrentTime
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CurrentTime{}
			if got := c.GetCurrentDayAndTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CurrentTime.GetCurrentDayAndTime() = %v, want %v", got, tt.want)
			}
		})
	}
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
}

func Test_cleanInput(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanInput(tt.args.inputString); got != tt.want {
				t.Errorf("cleanInput(%v) = %v, want %v", tt.args.inputString, got, tt.want)
			}
		})
	}
}

func TestStart(t *testing.T) {
	type args struct {
		templatePath  string
		outputDirPath string
		docType       string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Start(tt.args.templatePath, tt.args.outputDirPath, tt.args.docType)
		})
	}
}
