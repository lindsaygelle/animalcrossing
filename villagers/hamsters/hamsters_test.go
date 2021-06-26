package hamsters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllHamsters(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Hamster.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Hamster.Name())
		}
	}
}
