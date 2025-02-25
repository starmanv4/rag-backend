package rag

import (
	"strings"
)

// Simple keyword-based retrieval
func RetrieveContext(document, query string) string {
	lines := strings.Split(document, "\n")
	var relevant []string

	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), strings.ToLower(query)) {
			relevant = append(relevant, line)
		}
	}

	if len(relevant) == 0 {
		return "No relevant information found."
	}
	return strings.Join(relevant, "\n")
}
