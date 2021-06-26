package otters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLyleAnimal(t *testing.T) {
	var s string = animals.Otter.Name()
	if ok := Lyle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lyle.Animal(), s)
	}
}

func TestLyleName(t *testing.T) {
	if ok := Lyle.Name() == lyle; !ok {
		t.Fatalf("%s != %s", Lyle.Name(), lyle)
	}
}
