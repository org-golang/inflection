package inflection

type inflection interface {
	inflect(word string) string
}
