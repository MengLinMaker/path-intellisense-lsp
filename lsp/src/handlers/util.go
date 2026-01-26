package handlers

import (
	"errors"
	"os/user"
	"path/filepath"
	"regexp"
)

const (
	triggerCharacter   = "(\"|'|`| |\n)" // """ or "'" or "`" or " " or "\n"
	optionalPathPrefix = "([.]{1,2}|~)?" // "." or ".." or "~"
	illegalCharacters  = "\\/:?\"<>|\r\n&"
)

var regexCache = map[string]*regexp.Regexp{}

func mustCompileLazyRegex(filter string) *regexp.Regexp {
	if regexCache[filter] != nil {
		return regexCache[filter]
	}
	re := regexp.MustCompile(filter)
	regexCache[filter] = re
	return re
}

// Get last match of valid file path
func extractPathsRegex(text string) ([]string, error) {
	paths := []string{}

	re := mustCompileLazyRegex(triggerCharacter + optionalPathPrefix + "(/[^" + illegalCharacters + "]+)*/")
	matches := re.FindAllString("\n"+text, -1)
	if len(matches) == 0 {
		return paths, errors.New("no path matching strings found")
	}

	for _, match := range matches {
		paths = append(paths, match[1:])
	}
	return paths, nil
}

// Split text into lines
func textLines(text string) []string {
	return mustCompileLazyRegex("\r?\n").Split(text, -1)
}

func matchPath(path string, fileUri string, joinPath string) []string {
	switch string(path[0]) {
	case "/":
		return absolutePathSuggestions(path, joinPath)
	case "~":
		return homePathSuggestions(path, joinPath)
	case ".":
		return relativePathSuggestions(path, fileUri, joinPath)
	}
	return []string{}
}

func absolutePathSuggestions(absolutePath string, joinPath string) []string {
	searchPath := filepath.Join(absolutePath, joinPath)
	suggestedAbsolutePaths, err := filepath.Glob(searchPath)
	if err != nil {
		return []string{}
	}
	return suggestedAbsolutePaths
}

func homePathSuggestions(path string, joinPath string) []string {
	currentUser, err := user.Current()
	if err != nil {
		return []string{}
	}
	absolutePath := filepath.Join(currentUser.HomeDir, path[2:])
	return absolutePathSuggestions(absolutePath, joinPath)
}

func relativePathSuggestions(path string, fileUri string, joinPath string) []string {
	currentAbsoluteDirPath, _ := filepath.Split(fileUri[7:])
	absolutePath := filepath.Join(currentAbsoluteDirPath, path)
	return absolutePathSuggestions(absolutePath, joinPath)
}
