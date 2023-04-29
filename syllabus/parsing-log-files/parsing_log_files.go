package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DGB|INF|WRN|ERR|FTL)\].+`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`^.*".*((?i)password).*".*$`)
	for _, v := range lines {
		match := re.FindString(v)
		if len(match) > 0 {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`(?:User\s+)(\w+)`)

	for i, v := range lines {
		username := re.FindAllStringSubmatch(v, -1)
		if len(username) > 0 {
			lines[i] = "[USR] " + username[0][1] + " " + lines[i]
		}
	}

	return lines

}
