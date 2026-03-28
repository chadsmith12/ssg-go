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
