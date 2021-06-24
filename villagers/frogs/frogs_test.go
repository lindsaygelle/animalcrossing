package frogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllFrogs(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Frog.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Frog.Name())
		}
	}
}
