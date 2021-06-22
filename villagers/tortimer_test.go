package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTortimerAnimal(t *testing.T) {
	var s string = animals.Tortoise.Name()
	if ok := Tortimer.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tortimer.Animal(), s)
	}
}

func TestTortimerName(t *testing.T) {
	if ok := Tortimer.Name() == tortimer; !ok {
		t.Fatalf("%s != %s", Tortimer.Name(), tortimer)
	}
}
