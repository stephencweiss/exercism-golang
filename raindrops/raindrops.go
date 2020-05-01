// Package raindrops returns the sound made by a number of raindrops.
package raindrops

import "strconv"

// Convert is the function that takes an integer of rain drops and converts it to its noise equivalent.
// Noises are only emitted from raindrop counts that are factors of 3,5, or 7.
func Convert(raindropCount int) string {
	noise := ""
	if 0 == raindropCount%3 {
		noise += "Pling"
	}
	if 0 == raindropCount%5 {
		noise += "Plang"
	}
	if 0 == raindropCount%7 {
		noise += "Plong"
	}
	if noise == "" {
		noise = strconv.Itoa(raindropCount)
	}
	return noise
}
