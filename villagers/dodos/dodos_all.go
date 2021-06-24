package dodos

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllDodos(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Dodo.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Dodo.Name())
		}
	}
}
