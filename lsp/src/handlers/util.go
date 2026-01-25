package handlers

import (
	"errors"
	"regexp"
)

// Get last match of valid file path
func extractPathsRegex(text string) ([]string, error) {
	paths := []string{}

	triggerCharacter := "(\"|'|`| |\n)"   // """ or "'" or "`" or " " or "\n"
	optionalPathPrefix := "([.]{1,2}|~)?" // "." or ".." or "~"
	illegalCharacters := "\\/:?\"<>|\r\n&"
	re := regexp.MustCompile(triggerCharacter + optionalPathPrefix + "(/[^" + illegalCharacters + "]+)*/")
	matches := re.FindAllString("\n"+text, -1)
	if len(matches) == 0 {
		return paths, errors.New("no path matching strings found")
	}

	for _, match := range matches {
		paths = append(paths, match[1:])
	}
	return paths, nil
}
