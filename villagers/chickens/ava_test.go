package chickens

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAvaAnimal(t *testing.T) {
	var s string = animals.Chicken.Name()
	if ok := Ava.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ava.Animal(), s)
	}
}

func TestAvaName(t *testing.T) {
	if ok := Ava.Name() == ava; !ok {
		t.Fatalf("%s != %s", Ava.Name(), ava)
	}
}
