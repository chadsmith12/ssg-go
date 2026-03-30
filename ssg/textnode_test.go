package ssg_test

import (
	"testing"

	"github.com/chadsmith12/ssg-go/ssg"
)

func assertNodesEqualsFunc(t *testing.T, expected, got []*ssg.TextNode) {
	if len(expected) != len(got) {
		t.Errorf("len(nodes) wasn't %d, got %d", len(expected), len(got))
	}
	for i, node := range got {
		if !expected[i].Equals(node) {
			t.Errorf("node[%d] != expected[%d] wanted: %v - got %v",
				i, i, expected[i].String(), node.String())
		}
	}
}
func TestTextNodeEquals(t *testing.T) {
	assertEqualsFunc := func(t *testing.T, a, b *ssg.TextNode, want bool) {
		t.Helper()
		actual := a.Equals(b)

		if actual != want {
			textEquals := a.Text == b.Text
			ttEquals := a.TextType == b.TextType
			urlEquals := a.Url == b.Url
			t.Errorf("Equals() = %v, want %v  (Text: %v, TextType: %v, Url: %v)",
				actual, want, textEquals, ttEquals, urlEquals)
		}
	}

	t.Run("identical nodes are equal", func(t *testing.T) {
		node := ssg.CreateTextNode("This is a text node")
		node2 := ssg.CreateTextNode("This is a text node")
		assertEqualsFunc(t, node, node2, true)
	})

	t.Run("different text are not equal", func(t *testing.T) {
		node := ssg.CreateTextNode("This is a text node")
		node2 := ssg.CreateTextNode("This is also a text node")

		assertEqualsFunc(t, node, node2, false)
	})

	t.Run("different text types are not equal", func(t *testing.T) {
		node := ssg.CreateTextNode("This is a text node")
		node2 := ssg.CreateItalicTextNode("This is a text node")

		assertEqualsFunc(t, node, node2, false)
	})
}

func TestTextNodeDeliminator(t *testing.T) {

	t.Run("can create text nodes from single code block", func(t *testing.T) {
		node := ssg.CreateTextNode("This text is a `code block` word")
		nodes, err := node.SplitNodeDeliminator("`", ssg.TT_CODE)
		if err != nil {
			t.Errorf("err was not nil, got %v", err)
		}
		expected := []*ssg.TextNode{
			ssg.CreateTextNode("This text is a "),
			ssg.CreateCodeTextNode("code block"),
			ssg.CreateTextNode(" word"),
		}
		assertNodesEqualsFunc(t, expected, nodes)
	})

	t.Run("can create text nodes from bold block", func(t *testing.T) {
		node := ssg.CreateTextNode("This text is a **bold block** word")
		nodes, err := node.SplitNodeDeliminator("**", ssg.TT_BOLD)
		if err != nil {
			t.Errorf("err was not nil, got %v", err)
		}
		expected := []*ssg.TextNode{
			ssg.CreateTextNode("This text is a "),
			ssg.CreateBoldTextNode("bold block"),
			ssg.CreateTextNode(" word"),
		}
		assertNodesEqualsFunc(t, expected, nodes)
	})

	t.Run("can create text nodes from italic block", func(t *testing.T) {
		node := ssg.CreateTextNode("This text is a _italic block_ word")
		nodes, err := node.SplitNodeDeliminator("_", ssg.TT_ITALIC)
		if err != nil {
			t.Errorf("err was not nil, got %v", err)
		}

		expected := []*ssg.TextNode{
			ssg.CreateTextNode("This text is a "),
			ssg.CreateItalicTextNode("italic block"),
			ssg.CreateTextNode(" word"),
		}
		assertNodesEqualsFunc(t, expected, nodes)
	})

	t.Run("will return error with no ending deliminator", func(t *testing.T) {
		node := ssg.CreateTextNode("This text is a _italic block word")
		_, err := node.SplitNodeDeliminator("_", ssg.TT_ITALIC)
		if err == nil {
			t.Error("err was nil, expected NoEndingDeliminatorErr")
		}
	})
}

func TestLinkNodeSplitting(t *testing.T) {
	node := ssg.CreateTextNode("this is my text node with a [link](https://google.com) stuff after and another [link](https://google.com) with text ending")
	foundNodes := node.SplitNodeLinks()
	expectedNodes := []*ssg.TextNode{
		ssg.CreateTextNode("this is my text node with a "),
		ssg.CreateLinkTextNode("link", "https://google.com"),
		ssg.CreateTextNode(" stuff after and another "),
		ssg.CreateLinkTextNode("link", "https://google.com"),
		ssg.CreateTextNode(" with text ending"),
	}

	assertNodesEqualsFunc(t, expectedNodes, foundNodes)
}
