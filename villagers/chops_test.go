package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChopsAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Chops.Animal() == s; !ok {
		t.Fatalf("%s != %s", Chops.Animal(), s)
	}
}

func TestChopsName(t *testing.T) {
	if ok := Chops.Name() == chops; !ok {
		t.Fatalf("%s != %s", Chops.Name(), chops)
	}
}
