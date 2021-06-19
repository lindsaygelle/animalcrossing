package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBiffAnimal(t *testing.T) {
	var s string = animals.Hippo.Name()
	if ok := Biff.Animal() == s; !ok {
		t.Fatalf("%s != %s", Biff.Animal(), s)
	}
}

func TestBiffName(t *testing.T) {
	if ok := Biff.Name() == biff; !ok {
		t.Fatalf("%s != %s", Biff.Name(), biff)
	}
}
