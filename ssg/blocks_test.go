package ssg_test

import (
	"strings"
	"testing"

	"github.com/chadsmith12/ssg-go/ssg"
)

func makeHeading(level int, text string) string {
	builder := strings.Builder{}
	for range level {
		builder.WriteByte('#')
	}
	builder.WriteByte(' ')
	builder.WriteString(text)

	return builder.String()
}

func assertBlocksEqual(t *testing.T, expected, got []string) {
	t.Helper()

	if len(expected) != len(got) {
		t.Errorf("len(blocks) wasn't %d, got %d", len(expected), len(got))
	}
	for i, block := range got {
		if expected[i] != block {
			t.Errorf("node[%d] != expected[%d] wanted: %v - got %v",
				i, i, expected[i], block)
		}
	}
}

func TestMarkdownToBlocks(t *testing.T) {
	t.Run("can parse multiline markdown 1", func(t *testing.T) {

		markdown := `# This is a heading

This is a paragraph of text. It has some **bold** and _italic_ words inside of it.

- This is the first list item in a list block
- This is a list item
- This is another list item
		`

		blocks := ssg.MarkdownToBlocks(markdown)
		expected := []string{
			"# This is a heading",
			"This is a paragraph of text. It has some **bold** and _italic_ words inside of it.",
			"- This is the first list item in a list block\n- This is a list item\n- This is another list item",
		}

		assertBlocksEqual(t, expected, blocks)
	})
}

func TestBlockToBlockType(t *testing.T) {
	t.Run("can detect headings 1-6", func(t *testing.T) {
		for i := range 6 {
			heading := makeHeading(i+1, "hello world")
			expected := ssg.BT_HEADING
			if ssg.BlockToBlockType(heading) != expected {
				t.Errorf("wanted %s - got %s", expected, ssg.BlockToBlockType(heading))
			}
		}
	})

	t.Run("headings must have space after #", func(t *testing.T) {
		heading := "##Hello World"
		blockType := ssg.BlockToBlockType(heading)
		if blockType == ssg.BT_HEADING {
			t.Errorf("%s should not be a heading - got %s", heading, blockType)
		}
	})

	t.Run("empty text should return paragraph", func(t *testing.T) {
		heading := ""
		blockType := ssg.BlockToBlockType(heading)
		if blockType != ssg.BT_PARAGRAPH {
			t.Errorf("%s should be a paragraph - got %s", heading, blockType)
		}
	})

	t.Run("can detect multi-line code block", func(t *testing.T) {
		block := "```\nprint('hello world')\n```"
		blockType := ssg.BlockToBlockType(block)
		if blockType != ssg.BT_CODE {
			t.Errorf("%s should be a code block - got %s", block, blockType)
		}
	})

	t.Run("can detect quote block", func(t *testing.T) {
		block := ">Hello World\n> Hello World Line 2"
		blockType := ssg.BlockToBlockType(block)
		if blockType != ssg.BT_QUOTE {
			t.Errorf("%s should be a quote block - got %s", block, blockType)
		}
	})

	t.Run("quote block needs text", func(t *testing.T) {
		block := ">\n> Hello World Line 2"
		blockType := ssg.BlockToBlockType(block)
		if blockType != ssg.BT_PARAGRAPH {
			t.Errorf("%s should be a paragraph block - got %s", block, blockType)
		}
	})

	t.Run("can detect unordered list block", func(t *testing.T) {
		block := "- Item 1\n- Item 2\n- Item 3"
		blockType := ssg.BlockToBlockType(block)
		if blockType != ssg.BT_UL {
			t.Errorf("%s should be a unordered list block - got %s", block, blockType)
		}
	})

	t.Run("unordered list block lines must start with - ", func(t *testing.T) {
		block := "-Item 1\n-Item 2\n-Item 3"
		blockType := ssg.BlockToBlockType(block)
		if blockType != ssg.BT_PARAGRAPH {
			t.Errorf("%s should be a paragraph block - got %s", block, blockType)
		}
	})

	t.Run("detects ordered list block lines", func(t *testing.T) {
		block := "1. Item 1\n2. Item 2\n3. Item 4"
		blockType := ssg.BlockToBlockType(block)
		if blockType != ssg.BT_OL {
			t.Errorf("%s should be ordered list = got %s", block, blockType)
		}
	})

	t.Run("ordered lists numbers must be sequential", func(t *testing.T) {
		block := "1. Item 1\n3. Item 2"
		blockType := ssg.BlockToBlockType(block)
		if blockType == ssg.BT_OL {
			t.Errorf("%s should not be ordered list - got %s", block, blockType)
		}
	})
}
