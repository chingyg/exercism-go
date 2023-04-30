package hamming

import (
	"errors"
)

// Returns the hammer distance between two strings of equal lengths
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("the two strings are not the same length")
	}

	hammerDistance := 0

	for i := range a {
		if a[i] != b[i] {
			hammerDistance++
		}
	}

	return hammerDistance, nil
}
