package tigers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestClaudiaAnimal(t *testing.T) {
	var s string = animals.Tiger.Name()
	if ok := Claudia.Animal() == s; !ok {
		t.Fatalf("%s != %s", Claudia.Animal(), s)
	}
}

func TestClaudiaName(t *testing.T) {
	if ok := Claudia.Name() == claudia; !ok {
		t.Fatalf("%s != %s", Claudia.Name(), claudia)
	}
}
