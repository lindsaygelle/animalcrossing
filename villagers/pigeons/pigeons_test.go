package pigeons

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllPigeons(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Pigeon.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Pigeon.Name())
		}
	}
}
