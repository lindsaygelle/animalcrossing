package frogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHuckAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Huck.Animal() == s; !ok {
		t.Fatalf("%s != %s", Huck.Animal(), s)
	}
}

func TestHuckName(t *testing.T) {
	if ok := Huck.Name() == huck; !ok {
		t.Fatalf("%s != %s", Huck.Name(), huck)
	}
}
