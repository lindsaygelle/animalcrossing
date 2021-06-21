package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPetuniaRhinocerosRhinoAnimal(t *testing.T) {
	var s string = animals.Rhinoceros.Name()
	if ok := PetuniaRhinoceros.Animal() == s; !ok {
		t.Fatalf("%s != %s", PetuniaRhinoceros.Animal(), s)
	}
}

func TestPetuniaRhinocerosName(t *testing.T) {
	if ok := PetuniaRhinoceros.Name() == petunia; !ok {
		t.Fatalf("%s != %s", PetuniaRhinoceros.Name(), petunia)
	}
}
