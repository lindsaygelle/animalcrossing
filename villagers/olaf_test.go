package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestOlafAnimal(t *testing.T) {
	var s string = animals.Anteater.Name()
	if ok := Olaf.Animal() == s; !ok {
		t.Fatalf("%s != %s", Olaf.Animal(), s)
	}
}

func TestOlafName(t *testing.T) {
	if ok := Olaf.Name() == olaf; !ok {
		t.Fatalf("%s != %s", Olaf.Name(), olaf)
	}
}
