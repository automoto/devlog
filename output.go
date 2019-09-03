package main

import "fmt"

func generateMd(questions []string, otherSections []string) string{
	out := ""
	out += "### Dev Log\n"
	out += fmt.Sprintf("*created: %s*", getCurrentDayAndTime())
	for _, q := range questions {
		out += fmt.Sprintf("\n##### %s\n", q)
	}
	for _, q := range otherSections {
		out += fmt.Sprintf("\n##### %s\n", q)
	}

	return out
}
