package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllCats(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Cat.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Cat.Name())
		}
	}
}
