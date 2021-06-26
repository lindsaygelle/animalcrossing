package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestShakiAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Shaki.Animal() == s; !ok {
		t.Fatalf("%s != %s", Shaki.Animal(), s)
	}
}

func TestShakiName(t *testing.T) {
	if ok := Shaki.Name() == shaki; !ok {
		t.Fatalf("%s != %s", Shaki.Name(), shaki)
	}
}
