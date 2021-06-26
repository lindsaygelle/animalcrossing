package anteaters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllAnteaters(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Anteater.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Anteater.Name())
		}
	}
}
