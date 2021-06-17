package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRaddleAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Raddle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Raddle.Animal(), s)
	}
}

func TestRaddleName(t *testing.T) {
	if ok := Raddle.Name() == raddle; !ok {
		t.Fatalf("%s != %s", Raddle.Name(), raddle)
	}
}
