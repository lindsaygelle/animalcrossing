package bears

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBaroldAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Barold.Animal() == s; !ok {
		t.Fatalf("%s != %s", Barold.Animal(), s)
	}
}

func TestBaroldName(t *testing.T) {
	if ok := Barold.Name() == barold; !ok {
		t.Fatalf("%s != %s", Barold.Name(), barold)
	}
}
