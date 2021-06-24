package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFelyneAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Felyne.Animal() == s; !ok {
		t.Fatalf("%s != %s", Felyne.Animal(), s)
	}
}

func TestFelyneName(t *testing.T) {
	if ok := Felyne.Name() == felyne; !ok {
		t.Fatalf("%s != %s", Felyne.Name(), felyne)
	}
}
