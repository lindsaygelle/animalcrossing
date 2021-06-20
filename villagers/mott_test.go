package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMottAnimal(t *testing.T) {
	var s string = animals.Lion.Name()
	if ok := Mott.Animal() == s; !ok {
		t.Fatalf("%s != %s", Mott.Animal(), s)
	}
}

func TestMottName(t *testing.T) {
	if ok := Mott.Name() == mott; !ok {
		t.Fatalf("%s != %s", Mott.Name(), mott)
	}
}
