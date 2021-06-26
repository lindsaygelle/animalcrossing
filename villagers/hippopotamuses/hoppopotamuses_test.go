package hippopotamuses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllHippopotamuses(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Hippopotamus.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Hippopotamus.Name())
		}
	}
}
