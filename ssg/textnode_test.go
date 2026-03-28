package ssg_test

import (
	"testing"

	"github.com/chadsmith12/ssg-go/ssg"
)

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
		node := ssg.CreateTextNode("This is a text node", ssg.TT_BOLD)
		node2 := ssg.CreateTextNode("This is a text node", ssg.TT_BOLD)
		assertEqualsFunc(t, node, node2, true)
	})

	t.Run("different text are not equal", func(t *testing.T) {
		node := ssg.CreateTextNode("This is a text node", ssg.TT_BOLD)
		node2 := ssg.CreateTextNode("This is also a text node", ssg.TT_BOLD)

		assertEqualsFunc(t, node, node2, false)
	})

	t.Run("different text types are not equal", func(t *testing.T) {
		node := ssg.CreateTextNode("This is a text node", ssg.TT_BOLD)
		node2 := ssg.CreateTextNode("This is also a text node", ssg.TT_ITALIC)

		assertEqualsFunc(t, node, node2, false)
	})
}
