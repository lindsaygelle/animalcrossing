package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllDogs(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Dog.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Dog.Name())
		}
	}
}
