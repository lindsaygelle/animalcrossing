package boars

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllBoars(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Boar.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Boar.Name())
		}
	}
}
