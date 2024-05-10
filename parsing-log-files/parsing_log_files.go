package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	reg := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	return reg.MatchString(text)
}

func SplitLogLine(text string) []string {
	reg := regexp.MustCompile(`<[~*-=]*>`)
	return reg.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	reg := regexp.MustCompile(`".*[p|P][a|A][s|S]{2}[w|W][o|O][r|R][d|D].*"`)
	total := 0
	for _, line := range lines {
		total += len(reg.FindAllString(line, -1))
	}
	return total
}

func RemoveEndOfLineText(text string) string {
	reg := regexp.MustCompile(`end-of-line[0-9]+`)
	return string(reg.ReplaceAll([]byte(text), []byte{}))
}

func TagWithUserName(lines []string) []string {
	updatedLines := []string{}
	userName := ""
	reg := regexp.MustCompile(`User\s+([a-zA-Z0-9]+)`)
	for _, line := range lines {
		substrings := reg.FindStringSubmatch(line)
		newLine := line
		if substrings != nil {
			fmt.Println(line)
			userName = substrings[1]
			newLine = fmt.Sprintf("[USR] %s %s", userName, line)
		}
		updatedLines = append(updatedLines, newLine)
	}
	return updatedLines
}
