package boars

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDaisyMaeAnimal(t *testing.T) {
	var s string = animals.Boar.Name()
	if ok := DaisyMae.Animal() == s; !ok {
		t.Fatalf("%s != %s", DaisyMae.Animal(), s)
	}
}

func TestDaisyMaeName(t *testing.T) {
	if ok := DaisyMae.Name() == daisyMae; !ok {
		t.Fatalf("%s != %s", DaisyMae.Name(), daisyMae)
	}
}
