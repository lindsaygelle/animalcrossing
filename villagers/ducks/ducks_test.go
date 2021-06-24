package ducks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllDucks(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Duck.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Duck.Name())
		}
	}
}
