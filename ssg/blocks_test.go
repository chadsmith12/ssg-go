package ssg_test

import (
	"testing"

	"github.com/chadsmith12/ssg-go/ssg"
)

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
