package ssg_test

import (
	"testing"

	"github.com/chadsmith12/ssg-go/ssg"
)

func TextExtractMarkdownImages(t *testing.T) {
	assertFoundMarkDownImages := func(t *testing.T, expected, got []ssg.MarkdownImage) {
		if len(expected) != len(got) {
			t.Errorf("expected %d images, got %d", len(expected), len(got))
		}
		for i, image := range got {
			if !expected[i].Equals(image) {
				t.Errorf("expected[%d] != got[%d], expected %s - got %s", i, i, expected[i], got[i])
			}
		}
	}

	t.Run("can find single markdown image", func(t *testing.T) {
		images := ssg.ExtractMarkdownImages("this is my ![alt text](/image.png)")
		expected := []ssg.MarkdownImage{{AltText: "alt text", ImageUrl: "/image.png"}}
		assertFoundMarkDownImages(t, expected, images)
	})

	t.Run("can find multiple markdown images", func(t *testing.T) {
		images := ssg.ExtractMarkdownImages("this is my ![alt text](/image.png) and ![alt text](/image.png)")
		expected := []ssg.MarkdownImage{
			{AltText: "alt text", ImageUrl: "/image.png"},
			{AltText: "alt text", ImageUrl: "/image.png"},
		}
		assertFoundMarkDownImages(t, expected, images)
	})

	t.Run("finds no markdown images with just text", func(t *testing.T) {
		images := ssg.ExtractMarkdownImages("this is my text with no images")
		expected := []ssg.MarkdownImage{}
		assertFoundMarkDownImages(t, expected, images)
	})

	t.Run("finds no markdown images with invalid markdown", func(t *testing.T) {
		images := ssg.ExtractMarkdownImages("this is my text with ![alt text] not markdown ()")
		expected := []ssg.MarkdownImage{}
		assertFoundMarkDownImages(t, expected, images)
	})

	t.Run("finds no markdown images with invalid markdown - 2", func(t *testing.T) {
		images := ssg.ExtractMarkdownImages("this is my text with ![alt text](testing....]")
		expected := []ssg.MarkdownImage{}
		assertFoundMarkDownImages(t, expected, images)
	})
}

func TestExtractMarkdownLinks(t *testing.T) {
	assertFoundMarkdownLinks := func(t *testing.T, expected, got []ssg.MarkdownLink) {
		if len(expected) != len(got) {
			t.Errorf("expected %d links, got %d", len(expected), len(got))
		}
		for i, image := range got {
			if !expected[i].Equals(image) {
				t.Errorf("expected[%d] != got[%d], expected %s - got %s", i, i, expected[i], got[i])
			}
		}
	}

	t.Run("can find single markdown link", func(t *testing.T) {
		links := ssg.ExtractMarkdownLinks("this is my [alt text](/image.png)")
		expected := []ssg.MarkdownLink{{Text: "alt text", Url: "/image.png"}}
		assertFoundMarkdownLinks(t, expected, links)
	})

	t.Run("can find multiple markdown links", func(t *testing.T) {
		links := ssg.ExtractMarkdownLinks("this is my [alt text](/image.png) and another [alt text](/image.png)")
		expected := []ssg.MarkdownLink{
			{Text: "alt text", Url: "/image.png"},
			{Text: "alt text", Url: "/image.png"},
		}
		assertFoundMarkdownLinks(t, expected, links)
	})

	t.Run("finds no markdown links", func(t *testing.T) {
		links := ssg.ExtractMarkdownLinks("this is my text with no links")
		expected := []ssg.MarkdownLink{}
		assertFoundMarkdownLinks(t, expected, links)
	})

	t.Run("finds no markdown links with invalid markdown", func(t *testing.T) {
		links := ssg.ExtractMarkdownLinks("this is my text [invalid (link)")
		expected := []ssg.MarkdownLink{}
		assertFoundMarkdownLinks(t, expected, links)
	})

	t.Run("finds no markdown links with invalid markdown - 2", func(t *testing.T) {
		links := ssg.ExtractMarkdownLinks("this is my text with [alt text](testing....]")
		expected := []ssg.MarkdownLink{}
		assertFoundMarkdownLinks(t, expected, links)
	})

	t.Run("find link with text that only has a link", func(t *testing.T) {
		links := ssg.ExtractMarkdownLinks("[alt text](testing....)")
		expected := []ssg.MarkdownLink{
			{Text: "alt text", Url: "testing...."},
		}
		assertFoundMarkdownLinks(t, expected, links)
	})
}
