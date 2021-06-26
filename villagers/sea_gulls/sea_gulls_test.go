package sea_gulls

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllSeaGulls(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.SeaGull.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.SeaGull.Name())
		}
	}
}
