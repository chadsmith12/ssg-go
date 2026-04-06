package ssg

import (
	"fmt"
	"strconv"
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
	} else if isClodeBlock(block) {
		return BT_CODE
	} else if isQuoteBlock(block) {
		return BT_QUOTE
	} else if isUnorderedList(block) {
		return BT_UL
	} else if isOrderedList(block) {
		return BT_OL
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

func isClodeBlock(block string) bool {
	if len(block) < 3 {
		return false
	}

	blockLines := strings.Split(block, "\n")
	if len(blockLines) <= 2 {
		return false
	}

	if !strings.HasPrefix(blockLines[0], "```") {
		return false
	}

	lastLine := blockLines[len(blockLines)-1]
	if !strings.HasPrefix(lastLine, "```") {
		return false
	}

	return true
}

func isQuoteBlock(block string) bool {
	if block == "" {
		return false
	}

	blockLines := strings.SplitSeq(block, "\n")

	for blockLine := range blockLines {
		if blockLine[0] != '>' || len(blockLine) < 2 {
			return false
		}
	}

	return true
}

func isUnorderedList(block string) bool {
	if block == "" {
		return false
	}

	blockLines := strings.SplitSeq(block, "\n")
	for blockLine := range blockLines {
		if !strings.HasPrefix(blockLine, "- ") {
			return false
		}
	}

	return true
}

func isOrderedList(block string) bool {
	if block == "" {
		return false
	}

	blockLines := strings.SplitSeq(block, "\n")

	orderedNumber := 1
	for blockLine := range blockLines {
		currentNumber, err := strconv.Atoi(fmt.Sprintf("%c", blockLine[0]))
		if err != nil || currentNumber != orderedNumber {
			return false
		}
		if !strings.HasPrefix(blockLine, fmt.Sprintf("%d. ", currentNumber)) {
			return false
		}
		orderedNumber++
	}

	return true
}
