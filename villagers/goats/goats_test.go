package goats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllGoats(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Goat.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Goat.Name())
		}
	}
}
