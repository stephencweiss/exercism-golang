// Package proverb to convert a list of elements into a proverbial warning
package proverb

import (
	"fmt"
)

// Proverb takes a list of elements for a rhyme and constructs a proverb from them.
func Proverb(rhyme []string) (poem []string) {
	for index, value := range rhyme {
		if len(rhyme)-1 == index {
			poem = append(poem, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))
		} else {
			poem = append(poem, fmt.Sprintf("For want of a %v the %v was lost.", value, rhyme[index+1]))
		}
	}
	return
}
