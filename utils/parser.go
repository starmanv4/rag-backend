package utils

import (
	"strings"
)

// ParseText removes unwanted whitespace from the document
func ParseText(text string) string {
	lines := strings.Split(text, "\n")
	var cleanLines []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			cleanLines = append(cleanLines, trimmed)
		}
	}

	return strings.Join(cleanLines, "\n")
}
