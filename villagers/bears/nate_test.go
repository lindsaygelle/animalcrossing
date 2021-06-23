package bears

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestNateAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Nate.Animal() == s; !ok {
		t.Fatalf("%s != %s", Nate.Animal(), s)
	}
}

func TestNateName(t *testing.T) {
	if ok := Nate.Name() == nate; !ok {
		t.Fatalf("%s != %s", Nate.Name(), nate)
	}
}
