package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKoharuAnimal(t *testing.T) {
	var s string = animals.Kangaroo.Name()
	if ok := Koharu.Animal() == s; !ok {
		t.Fatalf("%s != %s", Koharu.Animal(), s)
	}
}

func TestKoharuName(t *testing.T) {
	if ok := Koharu.Name() == koharu; !ok {
		t.Fatalf("%s != %s", Koharu.Name(), koharu)
	}
}
