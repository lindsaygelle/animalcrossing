package chickens

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBenedictAnimal(t *testing.T) {
	var s string = animals.Chicken.Name()
	if ok := Benedict.Animal() == s; !ok {
		t.Fatalf("%s != %s", Benedict.Animal(), s)
	}
}

func TestBenedictName(t *testing.T) {
	if ok := Benedict.Name() == benedict; !ok {
		t.Fatalf("%s != %s", Benedict.Name(), benedict)
	}
}
