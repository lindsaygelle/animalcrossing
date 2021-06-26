package tortoises

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllTortoises(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Tortoise.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Tortoise.Name())
		}
	}
}
