// Package dna is designed to count the nucleotides in a sequence
package dna

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is the the incoming string received by the function
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	fmt.Println(d)
	for _, c := range d {
		if _, present := h[c]; !present {
			return nil, fmt.Errorf("Invalid nucleotide! %q is not valid", c)
		}
		h[c]++

	}
	return h, nil
}
