package kappas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllKappas(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Kappa.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Kappa.Name())
		}
	}
}
