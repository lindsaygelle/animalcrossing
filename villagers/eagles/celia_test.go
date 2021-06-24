package eagles

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCeliaAnimal(t *testing.T) {
	var s string = animals.Eagle.Name()
	if ok := Celia.Animal() == s; !ok {
		t.Fatalf("%s != %s", Celia.Animal(), s)
	}
}

func TestCeliaName(t *testing.T) {
	if ok := Celia.Name() == celia; !ok {
		t.Fatalf("%s != %s", Celia.Name(), celia)
	}
}
