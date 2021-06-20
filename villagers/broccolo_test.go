package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBroccoloAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Broccolo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Broccolo.Animal(), s)
	}
}

func TestBroccoloName(t *testing.T) {
	if ok := Broccolo.Name() == broccolo; !ok {
		t.Fatalf("%s != %s", Broccolo.Name(), broccolo)
	}
}
