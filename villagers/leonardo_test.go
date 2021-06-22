package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLeonardoAnimal(t *testing.T) {
	var s string = animals.Tiger.Name()
	if ok := Leonardo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Leonardo.Animal(), s)
	}
}

func TestLeonardoName(t *testing.T) {
	if ok := Leonardo.Name() == leonardo; !ok {
		t.Fatalf("%s != %s", Leonardo.Name(), leonardo)
	}
}
