package tortoises

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCornimerAnimal(t *testing.T) {
	var s string = animals.Tortoise.Name()
	if ok := Cornimer.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cornimer.Animal(), s)
	}
}

func TestCornimerName(t *testing.T) {
	if ok := Cornimer.Name() == cornimer; !ok {
		t.Fatalf("%s != %s", Cornimer.Name(), cornimer)
	}
}
