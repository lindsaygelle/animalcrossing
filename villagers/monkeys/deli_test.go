package monkeys

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDeliAnimal(t *testing.T) {
	var s string = animals.Monkey.Name()
	if ok := Deli.Animal() == s; !ok {
		t.Fatalf("%s != %s", Deli.Animal(), s)
	}
}

func TestDeliName(t *testing.T) {
	if ok := Deli.Name() == deli; !ok {
		t.Fatalf("%s != %s", Deli.Name(), deli)
	}
}
