package horses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllHorses(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Horse.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Horse.Name())
		}
	}
}
