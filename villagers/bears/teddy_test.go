package bears

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTeddyAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Teddy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Teddy.Animal(), s)
	}
}

func TestTeddyName(t *testing.T) {
	if ok := Teddy.Name() == teddy; !ok {
		t.Fatalf("%s != %s", Teddy.Name(), teddy)
	}
}
