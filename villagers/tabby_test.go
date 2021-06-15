package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTabbyAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Tabby.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tabby.Animal(), s)
	}
}

func TestTabbyName(t *testing.T) {
	if ok := Tabby.Name() == tabby; !ok {
		t.Fatalf("%s != %s", Tabby.Name(), tabby)
	}
}
