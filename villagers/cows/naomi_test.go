package cows

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestNaomiAnimal(t *testing.T) {
	var s string = animals.Cow.Name()
	if ok := Naomi.Animal() == s; !ok {
		t.Fatalf("%s != %s", Naomi.Animal(), s)
	}
}

func TestNaomiName(t *testing.T) {
	if ok := Naomi.Name() == naomi; !ok {
		t.Fatalf("%s != %s", Naomi.Name(), naomi)
	}
}
