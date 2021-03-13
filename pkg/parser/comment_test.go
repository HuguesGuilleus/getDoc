// getDoc
// 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parser

import (
	"fmt"
)

func ExampleTextComment() {
	c := TextComment(`All human beings are born free and equal in dignity and rights. They are endowed with reason and conscience and should act towards one another in a spirit of brotherhood. https://www.un.org/en/universal-declaration-human-rights/ First Article `)
	fmt.Printf("c.Code: %t\n", c.Code)
	for _, c := range c.Content {
		fmt.Println(c.URL, c.Text)
	}
	// Output: c.Code: false
	// false All human beings are born free and equal in dignity and rights. They are endowed with reason and conscience and should act towards one another in a spirit of brotherhood.
	// true https://www.un.org/en/universal-declaration-human-rights/
	// false First Article
}
