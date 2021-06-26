package otters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllOtters(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Otter.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Otter.Name())
		}
	}
}
