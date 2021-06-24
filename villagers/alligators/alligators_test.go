package alligators

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllAlligators(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Alligator.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Alligator.Name())
		}
	}
}
