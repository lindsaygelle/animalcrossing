package wolves

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTarouAnimal(t *testing.T) {
	var s string = animals.Wolf.Name()
	if ok := Tarou.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tarou.Animal(), s)
	}
}

func TestTarouName(t *testing.T) {
	if ok := Tarou.Name() == tarou; !ok {
		t.Fatalf("%s != %s", Tarou.Name(), tarou)
	}
}
