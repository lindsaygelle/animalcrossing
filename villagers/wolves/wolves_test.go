package wolves

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllWolves(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Wolf.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Wolf.Name())
		}
	}
}
