package horses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCleoAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Cleo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cleo.Animal(), s)
	}
}

func TestCleoName(t *testing.T) {
	if ok := Cleo.Name() == cleo; !ok {
		t.Fatalf("%s != %s", Cleo.Name(), cleo)
	}
}
