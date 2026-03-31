package main

import (
	"github.com/chadsmith12/ssg-go/ssg"
)

func main() {
	test := "This is **text** with an _italic_ word and a `code block` and an ![obi wan image](https://i.imgur.com/fJRm4Vk.jpeg) and a [link](https://boot.dev)"
	ssg.ExtractTextNodes(test)
}
