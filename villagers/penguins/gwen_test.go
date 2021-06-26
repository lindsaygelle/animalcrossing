package penguins

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGwenAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Gwen.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gwen.Animal(), s)
	}
}

func TestGwenName(t *testing.T) {
	if ok := Gwen.Name() == gwen; !ok {
		t.Fatalf("%s != %s", Gwen.Name(), gwen)
	}
}
