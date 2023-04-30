// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// Returns a string depending on the name parameter. If the name is empty, will return "One for you, one for me.".
// If the name is not empty, will return "One for <NAME>, one for me."
func ShareWith(name string) string {

	var processedName string

	if len(name) == 0 {
		processedName = "you"
	} else {
		processedName = name
	}

	return fmt.Sprintf("One for %s, one for me.", processedName)
}
