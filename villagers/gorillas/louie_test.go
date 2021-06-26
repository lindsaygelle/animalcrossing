package gorillas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLouieAnimal(t *testing.T) {
	var s string = animals.Gorilla.Name()
	if ok := Louie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Louie.Animal(), s)
	}
}

func TestLouieName(t *testing.T) {
	if ok := Louie.Name() == louie; !ok {
		t.Fatalf("%s != %s", Louie.Name(), louie)
	}
}
