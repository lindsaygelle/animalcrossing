package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAnalogAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Analog.Animal() == s; !ok {
		t.Fatalf("%s != %s", Analog.Animal(), s)
	}
}

func TestAnalogName(t *testing.T) {
	if ok := Analog.Name() == analog; !ok {
		t.Fatalf("%s != %s", Analog.Name(), analog)
	}
}
