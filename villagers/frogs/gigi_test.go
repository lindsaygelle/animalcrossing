package frogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGigiAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Gigi.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gigi.Animal(), s)
	}
}

func TestGigiName(t *testing.T) {
	if ok := Gigi.Name() == gigi; !ok {
		t.Fatalf("%s != %s", Gigi.Name(), gigi)
	}
}
