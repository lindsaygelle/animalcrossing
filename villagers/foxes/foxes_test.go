package foxes

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllFoxs(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Fox.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Fox.Name())
		}
	}
}
