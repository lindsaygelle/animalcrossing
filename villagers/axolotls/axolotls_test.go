package axolotls

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllAxolotls(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Axolotl.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Axolotl.Name())
		}
	}
}
