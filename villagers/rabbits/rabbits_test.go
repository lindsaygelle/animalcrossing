package rabbits

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAllRabbits(t *testing.T) {
	for i := 0; i < len(All); i++ {
		var v = All[i]
		if ok := v.Animal() == animals.Rabbit.Name(); !ok {
			t.Fatalf("%s != %s", v.Animal(), animals.Rabbit.Name())
		}
	}
}
