package search

import "regexp"

func search(search_string string, substring *regexp.Regexp) string {
	return substring.FindString(search_string)
}
