package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWillowAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Willow.Animal() == s; !ok {
		t.Fatalf("%s != %s", Willow.Animal(), s)
	}
}

func TestWillowName(t *testing.T) {
	if ok := Willow.Name() == willow; !ok {
		t.Fatalf("%s != %s", Willow.Name(), willow)
	}
}
