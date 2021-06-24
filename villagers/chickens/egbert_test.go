package chickens

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestEgbertAnimal(t *testing.T) {
	var s string = animals.Chicken.Name()
	if ok := Egbert.Animal() == s; !ok {
		t.Fatalf("%s != %s", Egbert.Animal(), s)
	}
}

func TestEgbertName(t *testing.T) {
	if ok := Egbert.Name() == egbert; !ok {
		t.Fatalf("%s != %s", Egbert.Name(), egbert)
	}
}
