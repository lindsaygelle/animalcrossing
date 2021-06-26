package walruses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWendellAnimal(t *testing.T) {
	var s string = animals.Walrus.Name()
	if ok := Wendell.Animal() == s; !ok {
		t.Fatalf("%s != %s", Wendell.Animal(), s)
	}
}

func TestWendellName(t *testing.T) {
	if ok := Wendell.Name() == wendell; !ok {
		t.Fatalf("%s != %s", Wendell.Name(), wendell)
	}
}
