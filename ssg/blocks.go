package ssg

import "strings"

func MarkdownToBlocks(text string) []string {
	splitBlocks := strings.Split(text, "\n\n")
	blocks := make([]string, 0, len(splitBlocks))

	for _, block := range splitBlocks {
		if block != "" {
			blocks = append(blocks, strings.TrimSpace(block))
		}
	}

	return blocks
}
