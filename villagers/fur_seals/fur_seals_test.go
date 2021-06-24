package fur_seals

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllFurSeals(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.FurSeal.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.FurSeal.Name())
		}
	}
}
