package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMedliAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Medli.Animal() == s; !ok {
		t.Fatalf("%s != %s", Medli.Animal(), s)
	}
}

func TestMedliName(t *testing.T) {
	if ok := Medli.Name() == medli; !ok {
		t.Fatalf("%s != %s", Medli.Name(), medli)
	}
}
