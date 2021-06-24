package bulls

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllBulls(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Bull.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Bull.Name())
		}
	}
}
