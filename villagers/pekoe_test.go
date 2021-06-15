package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPekoeAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Pekoe.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pekoe.Animal(), s)
	}
}

func TestPekoeName(t *testing.T) {
	if ok := Pekoe.Name() == pekoe; !ok {
		t.Fatalf("%s != %s", Pekoe.Name(), pekoe)
	}
}
