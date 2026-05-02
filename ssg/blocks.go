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

func (b BlockType) HtmlNode(text string) *HtmlNode {
	blockNode := DefaultHtmlNode()
	rawText := b.text(text)
	blockNode.Tag = b.tag(text)
	blockNode.Children = b.htmlChildren(text)
	if b == BT_HEADING {
		blockNode.Value = rawText
	}
	// get text/value for the node

	return blockNode
}

func (b BlockType) text(text string) string {
	switch b {
	case BT_PARAGRAPH:
		return text
	case BT_CODE:
		return getCodeText(text)
	case BT_HEADING:
		return getHeading(text)
	case BT_OL:
		return getOlText(text)
	case BT_UL:
		return getUlText(text)
	case BT_QUOTE:
		return getQuoteText(text)
	}

	return text
}

func getQuoteText(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimPrefix(strings.TrimPrefix(line, ">"), " ")
	}
	return strings.Join(lines, "\n")
}

func getCodeText(text string) string {
	lines := strings.Split(text, "\n")
	if len(lines) <= 2 {
		return ""
	}
	codeLines := lines[1 : len(lines)-1]
	return strings.Join(codeLines, "\n")
}

func getOlText(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		parts := strings.SplitN(line, ". ", 2)
		if len(parts) > 1 {
			lines[i] = parts[1]
		}
	}
	return strings.Join(lines, "\n")
}

func getUlText(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimPrefix(strings.TrimPrefix(line, "* "), "- ")
	}
	return strings.Join(lines, "\n")
}

func (b BlockType) htmlChildren(text string) []*HtmlNode {
	if b == BT_CODE {
		cleanText := b.text(text)
		return []*HtmlNode{CreateHtmlNode("code", cleanText, []*HtmlNode{}, make(map[string]string))}
	}
	if b == BT_PARAGRAPH {
		textNodes := ExtractTextNodes(text)
		htmlNodes := make([]*HtmlNode, len(textNodes))
		for i, node := range textNodes {
			htmlNodes[i] = node.HtmlNode()
		}
		return htmlNodes
	}
	if b == BT_OL || b == BT_UL {
		cleanText := b.text(text)
		textNodes := strings.Split(cleanText, "\n")
		listNodes := make([]*HtmlNode, len(textNodes))
		for i, textNode := range textNodes {
			listNodes[i] = CreateHtmlNode("li", textNode, []*HtmlNode{}, map[string]string{})
		}
		return listNodes
	}
	if b == BT_QUOTE {
		cleanText := b.text(text)
		return []*HtmlNode{CreateHtmlNode("p", cleanText, []*HtmlNode{}, map[string]string{})}
	}

	return []*HtmlNode{}
}

func (b BlockType) tag(text string) string {
	switch b {
	case BT_CODE:
		return "pre"
	case BT_OL:
		return "ol"
	case BT_UL:
		return "ul"
	case BT_QUOTE:
		return "blockquote"
	case BT_HEADING:
		return headingLevelTag(text)
	case BT_PARAGRAPH:
		return "p"
	}

	return "p"
}

func headingLevelTag(block string) string {
	// already assumes that bocktype is a heading.
	for i, currentChar := range block {
		if currentChar != '#' {
			return fmt.Sprintf("h%d", i)
		}
	}

	return "h6"
}

func getHeading(block string) string {
	for i, char := range block {
		if char != '#' && char != ' ' {
			return block[i:]
		}
	}

	return block
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
		if !strings.HasPrefix(blockLine, "- ") && !strings.HasPrefix(blockLine, "* ") {
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
