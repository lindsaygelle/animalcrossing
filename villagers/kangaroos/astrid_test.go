package kangaroos

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAstridAnimal(t *testing.T) {
	var s string = animals.Kangaroo.Name()
	if ok := Astrid.Animal() == s; !ok {
		t.Fatalf("%s != %s", Astrid.Animal(), s)
	}
}

func TestAstridName(t *testing.T) {
	if ok := Astrid.Name() == astrid; !ok {
		t.Fatalf("%s != %s", Astrid.Name(), astrid)
	}
}
