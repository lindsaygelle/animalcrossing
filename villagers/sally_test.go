package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSallyAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Sally.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sally.Animal(), s)
	}
}

func TestSallyName(t *testing.T) {
	if ok := Sally.Name() == sally; !ok {
		t.Fatalf("%s != %s", Sally.Name(), sally)
	}
}
