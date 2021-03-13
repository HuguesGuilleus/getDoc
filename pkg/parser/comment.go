// getDoc
// 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parser

import (
	"regexp"
	"strings"
)

type Comment struct {
	Code    bool
	Content []CommentText
}

// Paragraph with an alteration of text and URL.
type CommentText struct {
	URL  bool
	Text string
}

var getURL = regexp.MustCompile(`(\w+:\S+)`)

// Find URL in the text and return a comment.
func TextComment(text string) (c Comment) {

	begin := 0
	for _, extract := range getURL.FindAllStringIndex(text, -1) {
		if extract[0] != 0 {
			c.Content = append(c.Content, CommentText{
				URL:  false,
				Text: text[begin:extract[0]],
			})
		}
		c.Content = append(c.Content, CommentText{
			URL:  true,
			Text: text[extract[0]:extract[1]],
		})
		begin = extract[1]
	}

	if begin != len(text)-1 {
		c.Content = append(c.Content, CommentText{false, text[begin:]})
	}

	for i := range c.Content {
		c.Content[i].Text = strings.TrimSpace(c.Content[i].Text)
	}

	return
}
