package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLeopoldAnimal(t *testing.T) {
	var s string = animals.Lion.Name()
	if ok := Leopold.Animal() == s; !ok {
		t.Fatalf("%s != %s", Leopold.Animal(), s)
	}
}

func TestLeopoldName(t *testing.T) {
	if ok := Leopold.Name() == leopold; !ok {
		t.Fatalf("%s != %s", Leopold.Name(), leopold)
	}
}
