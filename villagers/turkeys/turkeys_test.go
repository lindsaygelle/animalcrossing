package turkeys

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllTurkeys(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Turkey.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Turkey.Name())
		}
	}
}
