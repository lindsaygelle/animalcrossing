package mice

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBettinaAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Bettina.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bettina.Animal(), s)
	}
}

func TestBettinaName(t *testing.T) {
	if ok := Bettina.Name() == bettina; !ok {
		t.Fatalf("%s != %s", Bettina.Name(), bettina)
	}
}
