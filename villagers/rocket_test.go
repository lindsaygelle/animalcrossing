package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRocketAnimal(t *testing.T) {
	var s string = animals.Gorilla.Name()
	if ok := Rocket.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rocket.Animal(), s)
	}
}

func TestRocketName(t *testing.T) {
	if ok := Rocket.Name() == rocket; !ok {
		t.Fatalf("%s != %s", Rocket.Name(), rocket)
	}
}
