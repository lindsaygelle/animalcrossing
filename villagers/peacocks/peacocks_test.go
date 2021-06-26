package peacocks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllPeacocks(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Peacock.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Peacock.Name())
		}
	}
}
