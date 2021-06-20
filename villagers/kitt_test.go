package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKittAnimal(t *testing.T) {
	var s string = animals.Kangaroo.Name()
	if ok := Kitt.Animal() == s; !ok {
		t.Fatalf("%s != %s", Kitt.Animal(), s)
	}
}

func TestKittName(t *testing.T) {
	if ok := Kitt.Name() == kitt; !ok {
		t.Fatalf("%s != %s", Kitt.Name(), kitt)
	}
}
