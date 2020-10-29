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
		// TODO: Add test cases.
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
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validDocTypes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validDocTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_archive(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			archive()
		})
	}
}

func Test_isDocTypeValid(t *testing.T) {
	type args struct {
		docTypeInput string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isDocTypeValid(tt.args.docTypeInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("isDocTypeValid(%v) error = %v, wantErr %v", tt.args.docTypeInput, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isDocTypeValid(%v) = %v, want %v", tt.args.docTypeInput, got, tt.want)
			}
		})
	}
}

func Test_handleError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleError(tt.args.err)
		})
	}
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
