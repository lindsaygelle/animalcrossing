package raccoons

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllRaccoons(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Raccoon.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Raccoon.Name())
		}
	}
}
