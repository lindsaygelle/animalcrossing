package ostriches

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllOstrichs(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Ostrich.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Ostrich.Name())
		}
	}
}
