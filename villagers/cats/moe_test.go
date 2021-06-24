package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMoeAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Moe.Animal() == s; !ok {
		t.Fatalf("%s != %s", Moe.Animal(), s)
	}
}

func TestMoeName(t *testing.T) {
	if ok := Moe.Name() == moe; !ok {
		t.Fatalf("%s != %s", Moe.Name(), moe)
	}
}
