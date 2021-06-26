package sloths

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllSloth(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Sloth.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Sloth.Name())
		}
	}
}
