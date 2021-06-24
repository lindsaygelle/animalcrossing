package eagles

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllEagles(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Eagle.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Eagle.Name())
		}
	}
}
