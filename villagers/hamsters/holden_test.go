package hamsters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHoldenAnimal(t *testing.T) {
	var s string = animals.Hamster.Name()
	if ok := Holden.Animal() == s; !ok {
		t.Fatalf("%s != %s", Holden.Animal(), s)
	}
}

func TestHoldenName(t *testing.T) {
	if ok := Holden.Name() == holden; !ok {
		t.Fatalf("%s != %s", Holden.Name(), holden)
	}
}
