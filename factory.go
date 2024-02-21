package inflection

import (
	"inflection/contracts"
	"sync"
)

var (
	wordInflector contracts.WordInflector
	once          sync.Once
)

func MakeInflector() contracts.WordInflector {

	once.Do(func() {
		wordInflector = &inflector{
			wordInflector: &pluralizer{caches: make(map[string]string)},
			snakeCache:    make(map[string]map[string]string),
		}
	})

	return wordInflector
}
