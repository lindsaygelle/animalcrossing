package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCydAnimal(t *testing.T) {
	var s string = animals.Elephant.Name()
	if ok := Cyd.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cyd.Animal(), s)
	}
}

func TestCydName(t *testing.T) {
	if ok := Cyd.Name() == cyd; !ok {
		t.Fatalf("%s != %s", Cyd.Name(), cyd)
	}
}
