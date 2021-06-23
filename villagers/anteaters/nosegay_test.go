package anteaters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestNosegayAnimal(t *testing.T) {
	var s string = animals.Anteater.Name()
	if ok := Nosegay.Animal() == s; !ok {
		t.Fatalf("%s != %s", Nosegay.Animal(), s)
	}
}

func TestNosegayName(t *testing.T) {
	if ok := Nosegay.Name() == nosegay; !ok {
		t.Fatalf("%s != %s", Nosegay.Name(), nosegay)
	}
}
