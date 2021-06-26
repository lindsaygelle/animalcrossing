package tapirs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllTapirs(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Tapir.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Tapir.Name())
		}
	}
}
