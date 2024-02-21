package inflection

import (
	"regexp"
	"strings"
)

type inflector struct {
	wordInflector inflection
	snakeCache    map[string]map[string]string
}

func (i *inflector) Plural(value string) string {

	words := i.split(value)

	if len(words) == 0 {
		return value
	}

	lastIndex := len(words) - 1

	words[lastIndex] = i.wordInflector.inflect(words[lastIndex])

	return strings.Join(words, "")
}

func (i *inflector) split(value string) []string {
	return strings.Split(i.replace(value), ",")
}

func (i *inflector) Snake(value, delimiter string) string {
	key := value

	if words, ok := i.snakeCache[key]; ok {
		if snakeWord, ok := words[delimiter]; ok {
			return snakeWord
		}
	}

	if value != strings.ToLower(value) {

		value = strings.ReplaceAll(value, " ", "")

		value = strings.ToLower(i.replace(value))
	}

	i.snakeCache[key] = map[string]string{delimiter: value}

	return value
}

func (i *inflector) replace(value string) string {
	return regexp.MustCompile("(.)([A-Z])").ReplaceAllString(value, "${1}_${2}")
}
