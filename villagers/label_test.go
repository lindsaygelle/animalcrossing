package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLabelAnimal(t *testing.T) {
	var s string = animals.Hedgehog.Name()
	if ok := Label.Animal() == s; !ok {
		t.Fatalf("%s != %s", Label.Animal(), s)
	}
}

func TestLabelName(t *testing.T) {
	if ok := Label.Name() == label; !ok {
		t.Fatalf("%s != %s", Label.Name(), label)
	}
}
