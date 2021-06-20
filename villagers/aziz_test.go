package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAzizAnimal(t *testing.T) {
	var s string = animals.Lion.Name()
	if ok := Aziz.Animal() == s; !ok {
		t.Fatalf("%s != %s", Aziz.Animal(), s)
	}
}

func TestAzizName(t *testing.T) {
	if ok := Aziz.Name() == aziz; !ok {
		t.Fatalf("%s != %s", Aziz.Name(), aziz)
	}
}
