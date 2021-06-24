package alpacas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllAlpacas(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Alpaca.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Alpaca.Name())
		}
	}
}
