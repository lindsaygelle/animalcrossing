package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCyrusAnimal(t *testing.T) {
	var s string = animals.Alpaca.Name()
	if ok := Cyrus.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cyrus.Animal(), s)
	}
}

func TestCyrusName(t *testing.T) {
	if ok := Cyrus.Name() == cyrus; !ok {
		t.Fatalf("%s != %s", Cyrus.Name(), cyrus)
	}
}
