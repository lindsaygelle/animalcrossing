package deers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllDeers(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Deer.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Deer.Name())
		}
	}
}
