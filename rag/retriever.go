package rag

import (
	"sort"
	"strings"
)

func RetrieveContext(document, query string) string {
	lines := strings.Split(document, "\n")
	queryWords := strings.Fields(strings.ToLower(query))
	relevance := make(map[string]int)

	for _, line := range lines {
		lowerLine := strings.ToLower(line)
		count := 0

		for _, word := range queryWords {
			if strings.Contains(lowerLine, word) {
				count++
			}
		}

		if count > 0 {
			relevance[line] = count
		}
	}

	if len(relevance) == 0 {
		return "No relevant information found."
	}

	sortedLines := make([]string, 0, len(relevance))
	for line := range relevance {
		sortedLines = append(sortedLines, line)
	}

	sort.Slice(sortedLines, func(i, j int) bool {
		return relevance[sortedLines[i]] > relevance[sortedLines[j]]
	})

	return strings.Join(sortedLines, "\n")
}
