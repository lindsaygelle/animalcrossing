package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSunnyAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Sunny.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sunny.Animal(), s)
	}
}

func TestSunnyName(t *testing.T) {
	if ok := Sunny.Name() == sunny; !ok {
		t.Fatalf("%s != %s", Sunny.Name(), sunny)
	}
}
