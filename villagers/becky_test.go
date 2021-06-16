package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBeckyAnimal(t *testing.T) {
	var s string = animals.Chicken.Name()
	if ok := Becky.Animal() == s; !ok {
		t.Fatalf("%s != %s", Becky.Animal(), s)
	}
}

func TestBeckyName(t *testing.T) {
	if ok := Becky.Name() == becky; !ok {
		t.Fatalf("%s != %s", Becky.Name(), becky)
	}
}
