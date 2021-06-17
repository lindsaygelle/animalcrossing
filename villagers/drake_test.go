package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDrakeAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Drake.Animal() == s; !ok {
		t.Fatalf("%s != %s", Drake.Animal(), s)
	}
}

func TestDrakeName(t *testing.T) {
	if ok := Drake.Name() == drake; !ok {
		t.Fatalf("%s != %s", Drake.Name(), drake)
	}
}
