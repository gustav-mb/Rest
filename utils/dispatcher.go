package utils

import "regexp"

// RestPattern is a regex pattern used for grouping and matching CRUD operations split up in:
// 1. The method (GET, POST, PUT, DELETE).
// 2. The url (e.g., students/id/1).
// 3. The body (e.g., the JSON, used for creating or updating a specific resource).
const RestPattern = "^(?P<method>[A-Za-z]+) *(?P<url>[A-Za-z/0-9]+)? *(?P<body>\\{[A-Za-z0-9\":,\\s]+\\})?"

// RegexDispatcher is a struct used for regexes with a capture group.
// It has support for creating a map, mapping a capture group name to the substring of the text provided.
type RegexDispatcher struct {
	regex *regexp.Regexp // regex is the compiled regular expression used by this RegexDispatcher.
}

// FindAllMatches matches a text string to the regular expression associated with this RegexDispatcher.
// Moreover, it creates a map, mapping a capture group name, to the actual substring match of the supplied text.
func (r *RegexDispatcher) FindAllMatches(text string) map[string]string {
	matches := r.regex.FindStringSubmatch(text)

	paramsMap := make(map[string]string)
	for i, name := range r.regex.SubexpNames() {
		if i > 0 && i <= len(matches) {
			paramsMap[name] = matches[i]
		}
	}

	return paramsMap
}

// NewRegexDispatcher creates and returns a new RegexDispatcher which will create a regular expression to match the specified pattern.
func NewRegexDispatcher(pattern string) *RegexDispatcher {
	return &RegexDispatcher{
		regex: regexp.MustCompile(pattern),
	}
}
