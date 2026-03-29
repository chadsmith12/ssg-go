package ssg

import (
	"fmt"
	"strings"
)

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

func (t *TextNode) SplitNodeDeliminator(deliminator string, textType TextType) []*TextNode {
	if t.TextType != TT_TEXT {
		return []*TextNode{t}
	}

	/**
	 * Loop through the text in the node one-by-one
		* start in search mode
		* if didn't find deliminator add char to builder
		* if found deliminator, skip deliminator length -1 (ie: `->1-1 =0, **->2-1=1)
		* if found deliminator, write to a new TextNode, add it, skip, continue on
		* in found mode:
		* if not deliminator add char to deliminator builder
		* if found deliminator, skip deliminator length -1
	*/

	nodes := make([]*TextNode, 0, 5)
	textBuilder := strings.Builder{}
	deliminatorBuilder := strings.Builder{}
	found := false
	for i := 0; i < len(t.Text); i++ {
		if t.Text[i] != deliminator[0] {
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

	nodes = append(nodes, CreateTextNode(textBuilder.String()))
	return nodes
}
