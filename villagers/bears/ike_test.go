package bears

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestIkeAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Ike.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ike.Animal(), s)
	}
}

func TestIkeName(t *testing.T) {
	if ok := Ike.Name() == ike; !ok {
		t.Fatalf("%s != %s", Ike.Name(), ike)
	}
}
