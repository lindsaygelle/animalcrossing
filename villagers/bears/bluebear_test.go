package bears

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBluebearAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Bluebear.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bluebear.Animal(), s)
	}
}

func TestBluebearName(t *testing.T) {
	if ok := Bluebear.Name() == bluebear; !ok {
		t.Fatalf("%s != %s", Bluebear.Name(), bluebear)
	}
}
