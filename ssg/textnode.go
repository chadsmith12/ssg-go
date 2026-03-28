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
