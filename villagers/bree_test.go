package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBreeAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Bree.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bree.Animal(), s)
	}
}

func TestBreeName(t *testing.T) {
	if ok := Bree.Name() == bree; !ok {
		t.Fatalf("%s != %s", Bree.Name(), bree)
	}
}
