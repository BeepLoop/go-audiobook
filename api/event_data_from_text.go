package api

import "regexp"

func DataFromTextSSE(eventString string) []string {
	re := regexp.MustCompile(`\{([^}]*)\}`)
	matches := re.FindAllStringSubmatch(eventString, -1)

	var result []string
	for _, match := range matches {
		if len(match) > 1 {
			result = append(result, match[1])
		}
	}

	return result
}
