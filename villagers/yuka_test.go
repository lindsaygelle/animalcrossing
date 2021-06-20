package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestYukaAnimal(t *testing.T) {
	var s string = animals.Koala.Name()
	if ok := Yuka.Animal() == s; !ok {
		t.Fatalf("%s != %s", Yuka.Animal(), s)
	}
}

func TestYukaName(t *testing.T) {
	if ok := Yuka.Name() == yuka; !ok {
		t.Fatalf("%s != %s", Yuka.Name(), yuka)
	}
}
