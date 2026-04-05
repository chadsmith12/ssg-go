package ssg

import (
	"strings"
)

type BlockType string

const (
	BT_PARAGRAPH BlockType = "paragraph"
	BT_HEADING   BlockType = "heading"
	BT_CODE      BlockType = "code"
	BT_QUOTE     BlockType = "quote"
	BT_UL        BlockType = "ul"
	BT_OL        BlockType = "ol"
)

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

func BlockToBlockType(block string) BlockType {
	if len(block) == 0 {
		return BT_PARAGRAPH
	}

	if isHeading(block) {
		return BT_HEADING
	}
	return BT_PARAGRAPH
}

func isHeading(block string) bool {
	if block[0] != '#' {
		return false
	}

	for _, char := range block {
		if char != '#' && char != ' ' {
			return false
		} else if char != '#' && char == ' ' {
			return true
		}
	}

	return false
}
