package kangaroos

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMarcieAnimal(t *testing.T) {
	var s string = animals.Kangaroo.Name()
	if ok := Marcie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Marcie.Animal(), s)
	}
}

func TestMarcieName(t *testing.T) {
	if ok := Marcie.Name() == marcie; !ok {
		t.Fatalf("%s != %s", Marcie.Name(), marcie)
	}
}
