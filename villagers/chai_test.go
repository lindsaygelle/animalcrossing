package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChaiAnimal(t *testing.T) {
	var s string = animals.Elephant.Name()
	if ok := Chai.Animal() == s; !ok {
		t.Fatalf("%s != %s", Chai.Animal(), s)
	}
}

func TestChaiName(t *testing.T) {
	if ok := Chai.Name() == chai; !ok {
		t.Fatalf("%s != %s", Chai.Name(), chai)
	}
}
