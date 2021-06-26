package pumpkins

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllPumpkins(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Pumpkin.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Pumpkin.Name())
		}
	}
}
