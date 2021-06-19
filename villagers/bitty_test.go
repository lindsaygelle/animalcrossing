package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBittyAnimal(t *testing.T) {
	var s string = animals.Hippo.Name()
	if ok := Bitty.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bitty.Animal(), s)
	}
}

func TestBittyName(t *testing.T) {
	if ok := Bitty.Name() == bitty; !ok {
		t.Fatalf("%s != %s", Bitty.Name(), bitty)
	}
}
