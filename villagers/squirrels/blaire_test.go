package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBlaireAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Blaire.Animal() == s; !ok {
		t.Fatalf("%s != %s", Blaire.Animal(), s)
	}
}

func TestBlaireName(t *testing.T) {
	if ok := Blaire.Name() == blaire; !ok {
		t.Fatalf("%s != %s", Blaire.Name(), blaire)
	}
}
