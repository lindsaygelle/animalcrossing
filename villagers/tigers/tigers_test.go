package tigers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllTigers(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Tiger.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Tiger.Name())
		}
	}
}
