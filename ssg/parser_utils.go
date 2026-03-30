package ssg

import (
	"strings"
)

type markdownParsingState int

const (
	searchingState markdownParsingState = iota
	textState
	urlState
)

type markdownReference struct {
	text string
	url  string
}

func (m markdownReference) Image() MarkdownImage {
	return MarkdownImage{
		AltText:  m.text,
		ImageUrl: m.url,
	}
}

func (m markdownReference) Link() MarkdownLink {
	return MarkdownLink{
		Text: m.text,
		Url:  m.url,
	}
}

type MarkdownImage struct {
	AltText  string
	ImageUrl string
}

func (i MarkdownImage) Equals(b MarkdownImage) bool {
	return i.AltText == b.AltText && i.ImageUrl == b.ImageUrl
}

type MarkdownLink struct {
	Text string
	Url  string
}

func (l MarkdownLink) Equals(b MarkdownLink) bool {
	return l.Text == b.Text && l.Url == b.Url
}

func extractMarkdownReference(text string, openToken string) []markdownReference {
	parsingState := searchingState
	altTextBuilder := strings.Builder{}
	urlBuilder := strings.Builder{}
	refsFound := make([]markdownReference, 0, 5)
	for i := 0; i < len(text); i++ {
		if parsingState == searchingState {
			// have we found the opening token?
			if strings.HasPrefix(text[i:], openToken) {
				parsingState = textState
				i += len(openToken) - 1
			}
		} else if parsingState == textState {
			if text[i] != ']' {
				altTextBuilder.WriteByte(text[i])
				continue
			}
			if strings.HasPrefix(text[i+1:], "(") {
				parsingState = urlState
				i += 1
				continue
			} else {
				// next character doesnt start the url, not a valid image section, move on
				parsingState = searchingState
				altTextBuilder.Reset()
				urlBuilder.Reset()
			}
		} else {
			if text[i] != ')' {
				urlBuilder.WriteByte(text[i])
			} else {
				// found a complete reference
				refsFound = append(refsFound, markdownReference{
					text: altTextBuilder.String(),
					url:  urlBuilder.String(),
				})
				altTextBuilder.Reset()
				urlBuilder.Reset()
				parsingState = searchingState
			}
		}
	}

	return refsFound
}

func ExtractMarkDownImages(text string) []MarkdownImage {
	references := extractMarkdownReference(text, "![")
	imagesFound := make([]MarkdownImage, len(references))
	for i, reference := range references {
		imagesFound[i] = reference.Image()
	}

	return imagesFound
}

func ExtractMarkdownLinks(text string) []MarkdownLink {
	references := extractMarkdownReference(text, "[")
	linksFound := make([]MarkdownLink, len(references))
	for i, reference := range references {
		linksFound[i] = reference.Link()
	}

	return linksFound
}
