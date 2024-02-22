## Inflection

#### Install
```shell
go get github.com/org-golang/inflection
```

#### Run test
```shell
 go test -v github.com/org-golang/inflection/testing
```

#### Simple
```go
import "github.com/org-golang/inflection"

// Get inflector instance.

inflector := inflection.MakeInflector()

// Convert a given inflection to a serpentine shape.
word := inflector.Snake("AreaCity", "_") // area_city

// Convert a given inflection to a plural.
inflector.Plural(word) // area_cities
```