## Inflection

#### simple
```go
import "template-new/api/helloworld/v2/inflection"

// Get inflector instance.

inflector := inflection.MakeInflector()

// Convert a given inflection to a serpentine shape.
word := inflector.Snake("MissionOrder", "_") // mission_order

// Convert a given inflection to a plural.
inflector.Plural(word) // mission_orders

```