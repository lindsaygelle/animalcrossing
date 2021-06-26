package rhinoceroses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSpikeAnimal(t *testing.T) {
	var s string = animals.Rhinoceros.Name()
	if ok := Spike.Animal() == s; !ok {
		t.Fatalf("%s != %s", Spike.Animal(), s)
	}
}

func TestSpikeName(t *testing.T) {
	if ok := Spike.Name() == spike; !ok {
		t.Fatalf("%s != %s", Spike.Name(), spike)
	}
}
