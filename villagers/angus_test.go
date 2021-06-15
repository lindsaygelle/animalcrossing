package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAngusAnimal(t *testing.T) {
	var s string = animals.Bull.Name()
	if ok := Angus.Animal() == s; !ok {
		t.Fatalf("%s != %s", Angus.Animal(), s)
	}
}

func TestAngusName(t *testing.T) {
	if ok := Angus.Name() == angus; !ok {
		t.Fatalf("%s != %s", Angus.Name(), angus)
	}
}
