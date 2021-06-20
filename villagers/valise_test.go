package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestValiseAnimal(t *testing.T) {
	var s string = animals.Kangaroo.Name()
	if ok := Valise.Animal() == s; !ok {
		t.Fatalf("%s != %s", Valise.Animal(), s)
	}
}

func TestValiseName(t *testing.T) {
	if ok := Valise.Name() == valise; !ok {
		t.Fatalf("%s != %s", Valise.Name(), valise)
	}
}
