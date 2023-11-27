package strings

import (
	"strings"
)

func CleanURL(inputURL string) string {
	cleanedURL := strings.TrimPrefix(inputURL, "http://www.")
	cleanedURL = strings.TrimPrefix(cleanedURL, "https://www.")
	cleanedURL = strings.TrimPrefix(cleanedURL, "www.")
	return cleanedURL
}
