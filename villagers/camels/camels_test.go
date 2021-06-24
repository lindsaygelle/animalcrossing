package camels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllCamels(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Camel.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Camel.Name())
		}
	}
}
