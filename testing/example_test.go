package testing

import (
	"github.com/org-golang/inflection"
	"testing"
)

var wordInflector = inflection.MakeInflector()

func TestSnake(t *testing.T) {

	inflection.MakeInflector()
	result := wordInflector.Snake("AreaCity", "_")

	if result != "area_city" {
		t.Errorf("Value [%s] not equals area_city", result)
	}
}

func TestPlural(t *testing.T) {
	result := wordInflector.Plural("area_city")

	if result != "area_cities" {
		t.Errorf("Value [%s] not equals area_cities ", result)
	}
}
