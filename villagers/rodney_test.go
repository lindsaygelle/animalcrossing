package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRodneyAnimal(t *testing.T) {
	var s string = animals.Hamster.Name()
	if ok := Rodney.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rodney.Animal(), s)
	}
}

func TestRodneyName(t *testing.T) {
	if ok := Rodney.Name() == rodney; !ok {
		t.Fatalf("%s != %s", Rodney.Name(), rodney)
	}
}
