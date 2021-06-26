package reindeers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllReindeers(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Reindeer.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Reindeer.Name())
		}
	}
}
