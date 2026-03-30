package ssg

import (
	"errors"
	"fmt"
	"strings"
)

var NoEndingDeliminatorErr = errors.New("ending deliminator was not found")

type TextType string

const (
	TT_TEXT   TextType = "text"
	TT_ITALIC TextType = "italic"
	TT_BOLD   TextType = "bold"
	TT_CODE   TextType = "code"
	TT_LINK   TextType = "link"
	TT_IMAGE  TextType = "image"
)

type TextNode struct {
	Text     string
	TextType TextType
	Url      string
}

// Creates a TextNode that is just generic text
func CreateTextNode(text string) *TextNode {
	return createTextNode(text, TT_TEXT, "")
}

// / Creates a TextNode that has italic text
func CreateItalicTextNode(text string) *TextNode {
	return createTextNode(text, TT_ITALIC, "")
}

// / Creates a TextNode that has bold text
func CreateBoldTextNode(text string) *TextNode {
	return createTextNode(text, TT_BOLD, "")
}

// / Creates a TextNode that is for code text
func CreateCodeTextNode(text string) *TextNode {
	return createTextNode(text, TT_CODE, "")
}

// Creates a TextNode that is for links
func CreateLinkTextNode(text, url string) *TextNode {
	return createTextNode(text, TT_LINK, url)
}

// Creates a TextNode that is for images
func CreateImageTextNode(text, imgSrc string) *TextNode {
	return createTextNode(text, TT_IMAGE, imgSrc)
}

func createTextNode(text string, textType TextType, url string) *TextNode {
	return &TextNode{
		Text:     text,
		TextType: textType,
		Url:      url,
	}
}

func (t *TextNode) Equals(other *TextNode) bool {
	return t.Text == other.Text && t.TextType == other.TextType && t.Url == other.Url
}

func (t *TextNode) String() string {
	return fmt.Sprintf("TextNode(%s, %s, %s)", t.Text, t.TextType, t.Url)
}

func (t *TextNode) HtmlNode() *HtmlNode {
	switch t.TextType {
	case TT_TEXT:
		return LeafHtmlNode("", t.Text)
	case TT_BOLD:
		return LeafHtmlNode("b", t.Text)
	case TT_CODE:
		return LeafHtmlNode("pre", t.Text)
	case TT_IMAGE:
		return CreateHtmlNode("img", "", nil, map[string]string{
			"src": t.Url,
			"alt": t.Text,
		})
	case TT_ITALIC:
		return LeafHtmlNode("i", t.Text)
	case TT_LINK:
		return CreateHtmlNode("a", t.Text, nil, map[string]string{
			"href": t.Url,
		})
	default:
		return LeafHtmlNode("", t.Text)
	}
}

func (t *TextNode) SplitNodeDeliminator(deliminator string, textType TextType) ([]*TextNode, error) {
	if t.TextType != TT_TEXT {
		return []*TextNode{t}, nil
	}

	nodes := make([]*TextNode, 0, 5)
	textBuilder := strings.Builder{}
	deliminatorBuilder := strings.Builder{}
	found := false
	for i := 0; i < len(t.Text); i++ {
		if !strings.HasPrefix(t.Text[i:], deliminator) {
			if !found {
				textBuilder.WriteByte(t.Text[i])
			} else {
				deliminatorBuilder.WriteByte(t.Text[i])
			}
		} else {
			if !found {
				nodes = append(nodes, CreateTextNode(textBuilder.String()))
				i += len(deliminator) - 1
				textBuilder.Reset()
			} else {
				nodes = append(nodes, createTextNode(deliminatorBuilder.String(), textType, ""))
				i += len(deliminator) - 1
				deliminatorBuilder.Reset()
			}
			found = !found
		}
	}

	// if we were in found mode we should probably throw an error as we didn't close the deliminator
	// otherwise write out the remaining text

	if found {
		return []*TextNode{}, NoEndingDeliminatorErr
	}
	nodes = append(nodes, CreateTextNode(textBuilder.String()))
	return nodes, nil
}

func (t *TextNode) SplitNodeLinks() []*TextNode {
	if t.TextType != TT_TEXT {
		return []*TextNode{}
	}

	links := ExtractMarkdownLinks(t.Text)
	if len(links) == 0 {
		return []*TextNode{t}
	}

	textNodes := make([]*TextNode, 0, 5)
	// go through all the links
	textToCut := t.Text
	for _, link := range links {
		rawText := fmt.Sprintf("[%s](%s)", link.Text, link.Url)
		before, after, _ := strings.Cut(textToCut, rawText)
		textNodes = append(textNodes, CreateTextNode(before))
		textNodes = append(textNodes, CreateLinkTextNode(link.Text, link.Url))
		textToCut = after
	}
	if textToCut != "" {
		textNodes = append(textNodes, CreateTextNode(textToCut))
	}

	return textNodes
}
