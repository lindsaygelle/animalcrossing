package bears

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllBears(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Bear.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Bear.Name())
		}
	}
}
