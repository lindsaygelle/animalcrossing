package beavers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllBeavers(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Beaver.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Beaver.Name())
		}
	}
}
