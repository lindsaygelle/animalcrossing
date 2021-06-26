package owls

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllOwls(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Owl.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Owl.Name())
		}
	}
}
