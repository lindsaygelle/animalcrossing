package giraffes

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllGiraffes(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Giraffe.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Giraffe.Name())
		}
	}
}
