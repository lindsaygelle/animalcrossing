package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPonchoAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Poncho.Animal() == s; !ok {
		t.Fatalf("%s != %s", Poncho.Animal(), s)
	}
}

func TestPonchoName(t *testing.T) {
	if ok := Poncho.Name() == poncho; !ok {
		t.Fatalf("%s != %s", Poncho.Name(), poncho)
	}
}
