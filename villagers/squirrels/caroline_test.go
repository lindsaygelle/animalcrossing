package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCarolineAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Caroline.Animal() == s; !ok {
		t.Fatalf("%s != %s", Caroline.Animal(), s)
	}
}

func TestCarolineName(t *testing.T) {
	if ok := Caroline.Name() == caroline; !ok {
		t.Fatalf("%s != %s", Caroline.Name(), caroline)
	}
}
