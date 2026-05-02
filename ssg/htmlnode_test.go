package ssg_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/chadsmith12/ssg-go/ssg"
)

func TestPropsToHtml(t *testing.T) {
	t.Run("empty props return empty string", func(t *testing.T) {
		node := ssg.DefaultHtmlNode()
		if got := node.Attributes(); got != "" {
			t.Errorf("expected empty string, got %q", got)
		}
	})

	t.Run("single props formats correctly", func(t *testing.T) {
		props := map[string]string{"href": "https://google.com"}
		node := ssg.CreateHtmlNode("a", "", nil, props)
		want := ` href="https://google.com"`

		if got := node.Attributes(); got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})

	t.Run("multiple props formats correctly", func(t *testing.T) {
		props := map[string]string{
			"href":   "https://google.com",
			"target": "_blank",
		}
		node := ssg.CreateHtmlNode("a", "", nil, props)

		got := node.Attributes()
		for attr, val := range node.Props {
			expected := fmt.Sprintf(` %s="%s"`, attr, val)
			if !strings.Contains(got, expected) {
				t.Errorf("output %q missing attribute %q", got, expected)
			}
		}
	})
}

func TestHtmlString(t *testing.T) {
	t.Run("can render leaf node with no attributes", func(t *testing.T) {
		node := ssg.LeafHtmlNode("p", "Hello World")
		htmlString := node.HtmlString()
		expected := "<p>Hello World</p>"
		if expected != htmlString {
			t.Errorf("expected %q, got %q", expected, htmlString)
		}
	})

	t.Run("can render leaf node with attributes", func(t *testing.T) {
		props := map[string]string{
			"href":   "https://google.com",
			"target": "_blank",
		}
		node := ssg.CreateHtmlNode("a", "", nil, props)

		got := node.HtmlString()
		for attr, val := range node.Props {
			expected := fmt.Sprintf(` %s="%s"`, attr, val)
			if !strings.Contains(got, expected) {
				t.Errorf("output %q missing attribute %q", got, expected)
			}
		}
		if !strings.Contains(got, "<a") {
			t.Errorf("output %q has wrong tag - expected %q", got, "<a")
		}
	})

	t.Run("can render just text html nodes", func(t *testing.T) {
		node := ssg.LeafHtmlNode("", "this is my text")
		expected := "this is my text"
		if got := node.HtmlString(); got != expected {
			t.Errorf("expected %q, got %q", "this is my text", got)
		}
	})

	t.Run("can render multiple html nodes", func(t *testing.T) {
		textNode := ssg.LeafHtmlNode("", "Hello World")
		bNode := ssg.LeafHtmlNode("b", "This is Bold")
		block := ssg.CreateHtmlNode("p", "", []*ssg.HtmlNode{textNode, bNode}, nil)
		got := block.HtmlString()
		expected := `<p>Hello World<b>This is Bold</b></p>`
		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})

	t.Run("can render html nodes with grand children", func(t *testing.T) {
		bNode := ssg.LeafHtmlNode("b", "This is Bold")
		spanNode := ssg.CreateHtmlNode("span", "", []*ssg.HtmlNode{bNode}, nil)
		divNode := ssg.CreateHtmlNode("div", "", []*ssg.HtmlNode{spanNode}, nil)
		got := divNode.HtmlString()
		expected := `<div><span><b>This is Bold</b></span></div>`
		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})

	t.Run("can render html nodes with attributes", func(t *testing.T) {
		aNode := ssg.CreateHtmlNode("a", "google", nil, map[string]string{"href": "google.com"})
		spanNode := ssg.CreateHtmlNode("span", "", []*ssg.HtmlNode{aNode}, nil)
		divNode := ssg.CreateHtmlNode("div", "", []*ssg.HtmlNode{spanNode}, nil)
		got := divNode.HtmlString()
		expected := `<div><span><a href="google.com">google</a></span></div>`
		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})
}

func TestFromMarkdown(t *testing.T) {
	t.Run("parses paragraph", func(t *testing.T) {
		md := "Hello World"
		node := ssg.FromMarkdown(md)
		got := node.HtmlString()
		expected := "<div><p>Hello World</p></div>"
		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})

	t.Run("parses heading", func(t *testing.T) {
		md := "# Hello World"
		node := ssg.FromMarkdown(md)
		got := node.HtmlString()
		expected := "<div><h1>Hello World</h1></div>"
		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})

	t.Run("parses multiple blocks", func(t *testing.T) {
		md := "# Hello World\n\nParagraph here"
		node := ssg.FromMarkdown(md)
		got := node.HtmlString()
		expected := "<div><h1>Hello World</h1><p>Paragraph here</p></div>"
		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})

	t.Run("parses code block", func(t *testing.T) {
		md := "```\ncode here\n```"
		node := ssg.FromMarkdown(md)
		got := node.HtmlString()
		expected := "<div><pre><code>code here</code></pre></div>"
		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})

	t.Run("parses quote block", func(t *testing.T) {
		md := "> Quote text"
		node := ssg.FromMarkdown(md)
		got := node.HtmlString()
		expected := "<div><blockquote><p>Quote text</p></blockquote></div>"
		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})

	t.Run("parses unordered list", func(t *testing.T) {
		md := "- item 1\n- item 2"
		node := ssg.FromMarkdown(md)
		got := node.HtmlString()
		expected := "<div><ul><li>item 1</li><li>item 2</li></ul></div>"
		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})

	t.Run("parses ordered list", func(t *testing.T) {
		md := "1. item 1\n2. item 2"
		node := ssg.FromMarkdown(md)
		got := node.HtmlString()
		expected := "<div><ol><li>item 1</li><li>item 2</li></ol></div>"
		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})
}
