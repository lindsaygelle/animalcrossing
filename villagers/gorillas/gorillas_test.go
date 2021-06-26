package gorillas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllGorillas(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Gorilla.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Gorilla.Name())
		}
	}
}
