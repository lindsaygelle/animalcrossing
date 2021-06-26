package lions

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBudAnimal(t *testing.T) {
	var s string = animals.Lion.Name()
	if ok := Bud.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bud.Animal(), s)
	}
}

func TestBudName(t *testing.T) {
	if ok := Bud.Name() == bud; !ok {
		t.Fatalf("%s != %s", Bud.Name(), bud)
	}
}
