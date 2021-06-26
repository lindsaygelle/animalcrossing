package moles

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDonResettiAnimal(t *testing.T) {
	var s string = animals.Mole.Name()
	if ok := DonResetti.Animal() == s; !ok {
		t.Fatalf("%s != %s", DonResetti.Animal(), s)
	}
}

func TestDonResettiName(t *testing.T) {
	if ok := DonResetti.Name() == donResetti; !ok {
		t.Fatalf("%s != %s", DonResetti.Name(), donResetti)
	}
}
