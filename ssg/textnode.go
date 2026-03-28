package ssg

import "fmt"

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

func CreateUrlTextNode(text string, textType TextType, url string) *TextNode {
	return &TextNode{
		Text:     text,
		TextType: textType,
		Url:      url,
	}
}

func CreateTextNode(text string, textType TextType) *TextNode {
	return CreateUrlTextNode(text, textType, "")
}

func (t *TextNode) Equals(other *TextNode) bool {
	return t.Text == other.Text && t.TextType == other.TextType && t.Url == other.Url
}

func (t *TextNode) String() string {
	return fmt.Sprintf("TextNode(%s, %s, %s)", t.Text, t.TextType, t.Url)
}
