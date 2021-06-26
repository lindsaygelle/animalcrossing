package kangaroos

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllKangaroos(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Kangaroo.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Kangaroo.Name())
		}
	}
}
