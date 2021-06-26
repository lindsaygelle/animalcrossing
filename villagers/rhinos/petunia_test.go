package rhinoceroses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPetuniaAnimal(t *testing.T) {
	var s string = animals.Rhinoceros.Name()
	if ok := Petunia.Animal() == s; !ok {
		t.Fatalf("%s != %s", Petunia.Animal(), s)
	}
}

func TestPetuniaName(t *testing.T) {
	if ok := Petunia.Name() == petunia; !ok {
		t.Fatalf("%s != %s", Petunia.Name(), petunia)
	}
}
