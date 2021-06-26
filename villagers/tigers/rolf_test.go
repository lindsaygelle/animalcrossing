package tigers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRolfAnimal(t *testing.T) {
	var s string = animals.Tiger.Name()
	if ok := Rolf.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rolf.Animal(), s)
	}
}

func TestRolfName(t *testing.T) {
	if ok := Rolf.Name() == rolf; !ok {
		t.Fatalf("%s != %s", Rolf.Name(), rolf)
	}
}
