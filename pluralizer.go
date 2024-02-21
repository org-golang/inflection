package inflection

import (
	"strings"
)

type pluralizer struct {
	caches map[string]string
}

func (p *pluralizer) inflect(word string) string {

	if pluralWord, ok := p.caches[word]; ok {
		return pluralWord
	}

	pluralWord := p.replace(word)

	p.caches[word] = pluralWord

	return pluralWord
}

func (p *pluralizer) replace(word string) string {

	has := p.hasChars(word)

	part := word[:len(word)-1]

	switch {
	case strings.HasSuffix(word, "s") || strings.HasSuffix(word, "x") || strings.HasSuffix(word, "z") || strings.HasSuffix(word, "ch") || strings.HasSuffix(word, "sh"):
		return word + "es"
	case strings.HasSuffix(word, "y") && has:
		return part + "ies"
	case strings.HasSuffix(word, "f"):
		return part + "ves"
	case strings.HasSuffix(word, "fe"):
		return word[:len(word)-2] + "ves"
	case strings.HasSuffix(word, "o") && has:
		return word + "es"
	default:
		return word + "s"
	}
}

func (p *pluralizer) hasChars(word string) bool {
	return strings.ContainsRune("bcdfghjklmnpqrstvwxyzs", rune(word[len(word)-2]))
}
