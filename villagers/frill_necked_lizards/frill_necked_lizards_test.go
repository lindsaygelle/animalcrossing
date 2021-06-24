package frill_necked_lizards

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllFrillNeckedLizards(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.FrillNeckedLizard.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.FrillNeckedLizard.Name())
		}
	}
}
