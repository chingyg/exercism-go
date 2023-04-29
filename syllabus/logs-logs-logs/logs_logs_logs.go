package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {

	runeMap := map[rune]string{
		'\u2757':     "recommendation",
		'\U0001F50D': "search",
		'\u2600':     "weather"}

	for _, logValue := range log {
		if runeMapValue, exists := runeMap[logValue]; exists {
			return runeMapValue
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {

	var result []rune

	// No new replacements
	if oldRune == newRune {
		return log
	}
	for _, value := range log {

		if value == oldRune {
			result = append(result, newRune)

		} else {
			result = append(result, value)
		}
	}

	return string(result)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
