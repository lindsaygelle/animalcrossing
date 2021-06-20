package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWaltAnimal(t *testing.T) {
	var s string = animals.Kangaroo.Name()
	if ok := Walt.Animal() == s; !ok {
		t.Fatalf("%s != %s", Walt.Animal(), s)
	}
}

func TestWaltName(t *testing.T) {
	if ok := Walt.Name() == walt; !ok {
		t.Fatalf("%s != %s", Walt.Name(), walt)
	}
}
