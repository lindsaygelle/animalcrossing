package lions

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLionelAnimal(t *testing.T) {
	var s string = animals.Lion.Name()
	if ok := Lionel.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lionel.Animal(), s)
	}
}

func TestLionelName(t *testing.T) {
	if ok := Lionel.Name() == lionel; !ok {
		t.Fatalf("%s != %s", Lionel.Name(), lionel)
	}
}
