package pkg

import "fmt"

var Version = "v0.0.8"

func getVersion() string{
	return fmt.Sprintf(Version)
}
