package sheep

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllSheep(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Sheep.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Sheep.Name())
		}
	}
}
