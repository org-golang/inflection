package contracts

type WordInflector interface {

	// Plural Convert a given inflection to a plural.
	Plural(value string) string

	// Snake Convert a given inflection to a serpentine shape.
	Snake(value, delimiter string) string
}
