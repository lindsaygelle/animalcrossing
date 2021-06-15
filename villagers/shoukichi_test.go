package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestShoukichiAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Shoukichi.Animal() == s; !ok {
		t.Fatalf("%s != %s", Shoukichi.Animal(), s)
	}
}

func TestShoukichiName(t *testing.T) {
	if ok := Shoukichi.Name() == shoukichi; !ok {
		t.Fatalf("%s != %s", Shoukichi.Name(), shoukichi)
	}
}
