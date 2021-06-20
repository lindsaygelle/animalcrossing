package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTammiAnimal(t *testing.T) {
	var s string = animals.Monkey.Name()
	if ok := Tammi.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tammi.Animal(), s)
	}
}

func TestTammiName(t *testing.T) {
	if ok := Tammi.Name() == tammi; !ok {
		t.Fatalf("%s != %s", Tammi.Name(), tammi)
	}
}
