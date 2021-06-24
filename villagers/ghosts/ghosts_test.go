package ghosts

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllGhosts(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Ghost.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Ghost.Name())
		}
	}
}
