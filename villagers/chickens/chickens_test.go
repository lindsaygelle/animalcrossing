package chickens

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllChickens(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Chicken.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Chicken.Name())
		}
	}
}
