package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPhoebeAnimal(t *testing.T) {
	var s string = animals.Ostrich.Name()
	if ok := Phoebe.Animal() == s; !ok {
		t.Fatalf("%s != %s", Phoebe.Animal(), s)
	}
}

func TestPhoebeName(t *testing.T) {
	if ok := Phoebe.Name() == phoebe; !ok {
		t.Fatalf("%s != %s", Phoebe.Name(), phoebe)
	}
}
