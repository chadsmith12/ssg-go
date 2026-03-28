package ssg

import (
	"fmt"
	"strings"
)

type HtmlNodeRenderer interface {
	Html(*HtmlNode) string
}

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

// Attributes returns the props for this HTMLNode as HTML attributes string
func (n *HtmlNode) Attributes() string {
	if n.Props == nil {
		return ""
	}

	builder := strings.Builder{}
	for attr, value := range n.Props {
		fmt.Fprintf(&builder, " %s=\"%s\"", attr, value)
	}

	return builder.String()
}
