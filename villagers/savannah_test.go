package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSavannahAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Savannah.Animal() == s; !ok {
		t.Fatalf("%s != %s", Savannah.Animal(), s)
	}
}

func TestSavannahName(t *testing.T) {
	if ok := Savannah.Name() == savannah; !ok {
		t.Fatalf("%s != %s", Savannah.Name(), savannah)
	}
}
