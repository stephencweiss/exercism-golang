/*
Package hamming is used to calculate the hamming distance between two strings
For more on the hamming distance, see the wikipedia entry: https://en.wikipedia.org/wiki/Hamming_distance
*/
package hamming

import (
	"errors"
)

// Distance is a function that returns a hamming distance between two strings of equal length
// if the length of the strings is not equal, Distance throws an error
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("can't work with this")
	}
	distance := 0
	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			distance++
		}
	}
	return distance, nil
}
