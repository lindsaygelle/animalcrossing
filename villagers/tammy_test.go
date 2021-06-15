package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTammyAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Tammy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tammy.Animal(), s)
	}
}

func TestTammyName(t *testing.T) {
	if ok := Tammy.Name() == tammy; !ok {
		t.Fatalf("%s != %s", Tammy.Name(), tammy)
	}
}
