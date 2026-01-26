package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"os/user"
	"path/filepath"
	"regexp"
)

const (
	triggerCharacter   = "(\"|'|`| |\n)"          // """ or "'" or "`" or " " or "\n"
	optionalPathPrefix = "([.]{1,2}|~)?"          // "." or ".." or "~"
	illegalCharacters  = "\\/:?<>|\r" + "\"'` \n" // Must include trigger characters
)

var regexCache = map[string]*regexp.Regexp{}

// Lazy compile regex and cache for reuse
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
	re := mustCompileLazyRegex(triggerCharacter + optionalPathPrefix + fmt.Sprintf("(/[^%s]+)*", illegalCharacters) + "/")
	matches := re.FindAllString("\n"+text, -1)
	if len(matches) == 0 {
		return []string{}, errors.New("no path matching strings found")
	}
	paths := []string{}
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
	absolutePath := filepath.Join(currentUser.HomeDir, path[1:])
	return absolutePathSuggestions(absolutePath, joinPath)
}

func relativePathSuggestions(path string, fileUri string, joinPath string) []string {
	currentAbsoluteDirPath, _ := filepath.Split(fileUri[7:])
	absolutePath := filepath.Join(currentAbsoluteDirPath, path)
	return absolutePathSuggestions(absolutePath, joinPath)
}

type pathMatch struct {
	Text  string
	Start int
	End   int
}

func findPathMatches(line string) []pathMatch {
	re := mustCompileLazyRegex(triggerCharacter + optionalPathPrefix + fmt.Sprintf("(/[^%s]+)+", illegalCharacters))
	matches := re.FindAllStringIndex("\n"+line, -1)

	results := make([]pathMatch, 0, len(matches))
	for _, loc := range matches {
		// By spec len(loc) == 2
		start := loc[0]
		end := loc[1] - 1
		if start < 0 || end <= start || end > len(line) {
			slog.Error(fmt.Sprintf("Failed to extract path from line:\n%s\nstart(%d), end(%d), len(%d)", line, start, end, len(line)))
			continue
		}
		results = append(results, pathMatch{
			Start: start,
			End:   end,
			Text:  line[start:end],
		})
	}
	return results
}
