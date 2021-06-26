package chameleons

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllChameleons(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Chameleon.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Chameleon.Name())
		}
	}
}