package lions

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllLions(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Lion.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Lion.Name())
		}
	}
}