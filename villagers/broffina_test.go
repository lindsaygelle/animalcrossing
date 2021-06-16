package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBroffinaAnimal(t *testing.T) {
	var s string = animals.Chicken.Name()
	if ok := Broffina.Animal() == s; !ok {
		t.Fatalf("%s != %s", Broffina.Animal(), s)
	}
}

func TestBroffinaName(t *testing.T) {
	if ok := Broffina.Name() == broffina; !ok {
		t.Fatalf("%s != %s", Broffina.Name(), broffina)
	}
}
