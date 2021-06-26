package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllSquirrels(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Squirrel.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Squirrel.Name())
		}
	}
}
