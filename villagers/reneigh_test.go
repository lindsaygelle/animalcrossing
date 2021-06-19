package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestReneighAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Reneigh.Animal() == s; !ok {
		t.Fatalf("%s != %s", Reneigh.Animal(), s)
	}
}

func TestReneighName(t *testing.T) {
	if ok := Reneigh.Name() == reneigh; !ok {
		t.Fatalf("%s != %s", Reneigh.Name(), reneigh)
	}
}
