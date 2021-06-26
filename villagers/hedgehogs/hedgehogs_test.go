package hedgehogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllHedgehogs(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Hedgehog.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Hedgehog.Name())
		}
	}
}
