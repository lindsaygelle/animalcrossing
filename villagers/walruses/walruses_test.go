package walruses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllWalruses(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Walrus.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Walrus.Name())
		}
	}
}
