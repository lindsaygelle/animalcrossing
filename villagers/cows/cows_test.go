package cows

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllCows(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Cow.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Cow.Name())
		}
	}
}
