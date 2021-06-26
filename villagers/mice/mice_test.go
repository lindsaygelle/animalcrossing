package mice

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllMice(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Mouse.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Mouse.Name())
		}
	}
}
