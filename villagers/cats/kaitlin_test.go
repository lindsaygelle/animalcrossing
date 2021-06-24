package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKaitlinAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Kaitlin.Animal() == s; !ok {
		t.Fatalf("%s != %s", Kaitlin.Animal(), s)
	}
}

func TestKaitlinName(t *testing.T) {
	if ok := Kaitlin.Name() == kaitlin; !ok {
		t.Fatalf("%s != %s", Kaitlin.Name(), kaitlin)
	}
}
