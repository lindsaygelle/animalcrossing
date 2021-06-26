package rhinoceroses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllRhinoceross(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Rhinoceros.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Rhinoceros.Name())
		}
	}
}
