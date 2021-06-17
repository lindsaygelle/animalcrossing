package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestButchAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Butch.Animal() == s; !ok {
		t.Fatalf("%s != %s", Butch.Animal(), s)
	}
}

func TestButchName(t *testing.T) {
	if ok := Butch.Name() == butch; !ok {
		t.Fatalf("%s != %s", Butch.Name(), butch)
	}
}
