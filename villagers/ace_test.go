package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAceAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Ace.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ace.Animal(), s)
	}
}

func TestAceName(t *testing.T) {
	if ok := Ace.Name() == ace; !ok {
		t.Fatalf("%s != %s", Ace.Name(), ace)
	}
}