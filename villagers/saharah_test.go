package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSaharahAnimal(t *testing.T) {
	var s string = animals.Camel.Name()
	if ok := Saharah.Animal() == s; !ok {
		t.Fatalf("%s != %s", Saharah.Animal(), s)
	}
}

func TestSaharahName(t *testing.T) {
	if ok := Saharah.Name() == saharah; !ok {
		t.Fatalf("%s != %s", Saharah.Name(), saharah)
	}
}
