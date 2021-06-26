package penguins

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllPenguins(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Penguin.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Penguin.Name())
		}
	}
}
