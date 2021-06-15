package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAntonioAnimal(t *testing.T) {
	var s string = animals.Anteater.Name()
	if ok := Antonio.Animal() == s; !ok {
		t.Fatalf("%s != %s", Antonio.Animal(), s)
	}
}

func TestAntonioName(t *testing.T) {
	if ok := Antonio.Name() == antonio; !ok {
		t.Fatalf("%s != %s", Antonio.Name(), antonio)
	}
}
