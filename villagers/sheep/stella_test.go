package sheep

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestStellaAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Stella.Animal() == s; !ok {
		t.Fatalf("%s != %s", Stella.Animal(), s)
	}
}

func TestStellaName(t *testing.T) {
	if ok := Stella.Name() == stella; !ok {
		t.Fatalf("%s != %s", Stella.Name(), stella)
	}
}
