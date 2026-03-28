package ssg

import (
	"fmt"
	"strings"
)

type HtmlNode struct {
	Tag      string
	Value    string
	Children []*HtmlNode
	Props    map[string]string
}

func DefaultHtmlNode() *HtmlNode {
	return &HtmlNode{
		Children: make([]*HtmlNode, 0),
		Props:    map[string]string{},
	}
}

func CreateHtmlNode(tag, value string, children []*HtmlNode, props map[string]string) *HtmlNode {
	return &HtmlNode{
		Tag:      tag,
		Value:    value,
		Children: children,
		Props:    props,
	}
}

func LeafHtmlNode(tag, value string) *HtmlNode {
	return CreateHtmlNode(tag, value, []*HtmlNode{}, map[string]string{})
}

// Attributes returns the props for this HTMLNode as HTML attributes string
func (n *HtmlNode) Attributes() string {
	if len(n.Props) == 0 {
		return ""
	}

	builder := strings.Builder{}
	for attr, value := range n.Props {
		fmt.Fprintf(&builder, " %s=\"%s\"", attr, value)
	}

	return builder.String()
}

// builds and returns the HtmlNode as a html string
func (n *HtmlNode) HtmlString() string {
	builder := strings.Builder{}

	if len(n.Children) == 0 {
		fmt.Fprintf(&builder, `%s%s%s`, n.createOpeningTagString(), n.Value, n.createClosingTagString())
		return builder.String()
	}

	builder.WriteString(n.createOpeningTagString())
	for _, child := range n.Children {
		builder.WriteString(child.HtmlString())
	}
	builder.WriteString(n.createClosingTagString())

	return builder.String()
}

func (n *HtmlNode) createOpeningTagString() string {
	// not every html node has a tag, could just be text
	if n.Tag == "" {
		return ""
	}
	attrs := n.Attributes()
	if attrs == "" {
		return fmt.Sprintf("<%s>", n.Tag)
	}

	return fmt.Sprintf("<%s%s>", n.Tag, attrs)
}

func (n *HtmlNode) createClosingTagString() string {
	if n.Tag == "" {
		return ""
	}

	return fmt.Sprintf("</%s>", n.Tag)
}
