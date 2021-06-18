package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSherbAnimal(t *testing.T) {
	var s string = animals.Goat.Name()
	if ok := Sherb.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sherb.Animal(), s)
	}
}

func TestSherbName(t *testing.T) {
	if ok := Sherb.Name() == sherb; !ok {
		t.Fatalf("%s != %s", Sherb.Name(), sherb)
	}
}
