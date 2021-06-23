package bears

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMapleAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Maple.Animal() == s; !ok {
		t.Fatalf("%s != %s", Maple.Animal(), s)
	}
}

func TestMapleName(t *testing.T) {
	if ok := Maple.Name() == maple; !ok {
		t.Fatalf("%s != %s", Maple.Name(), maple)
	}
}
