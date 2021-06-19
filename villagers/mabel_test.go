package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMabelAnimal(t *testing.T) {
	var s string = animals.Hedgehog.Name()
	if ok := Mabel.Animal() == s; !ok {
		t.Fatalf("%s != %s", Mabel.Animal(), s)
	}
}

func TestMabelName(t *testing.T) {
	if ok := Mabel.Name() == mabel; !ok {
		t.Fatalf("%s != %s", Mabel.Name(), mabel)
	}
}
