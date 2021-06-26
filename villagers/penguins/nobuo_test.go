package penguins

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestNobuoAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Nobuo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Nobuo.Animal(), s)
	}
}

func TestNobuoName(t *testing.T) {
	if ok := Nobuo.Name() == nobuo; !ok {
		t.Fatalf("%s != %s", Nobuo.Name(), nobuo)
	}
}
