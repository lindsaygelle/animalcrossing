package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMasaAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Masa.Animal() == s; !ok {
		t.Fatalf("%s != %s", Masa.Animal(), s)
	}
}

func TestMasaName(t *testing.T) {
	if ok := Masa.Name() == masa; !ok {
		t.Fatalf("%s != %s", Masa.Name(), masa)
	}
}
