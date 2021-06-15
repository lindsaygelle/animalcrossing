package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJacquesAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Jacques.Animal() == s; !ok {
		t.Fatalf("%s != %s", Jacques.Animal(), s)
	}
}

func TestJacquesName(t *testing.T) {
	if ok := Jacques.Name() == jacques; !ok {
		t.Fatalf("%s != %s", Jacques.Name(), jacques)
	}
}
