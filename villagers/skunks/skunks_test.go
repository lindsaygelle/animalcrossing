package skunks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllSkunk(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Skunk.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Skunk.Name())
		}
	}
}
