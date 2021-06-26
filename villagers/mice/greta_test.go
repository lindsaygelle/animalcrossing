package mice

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGretaAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Greta.Animal() == s; !ok {
		t.Fatalf("%s != %s", Greta.Animal(), s)
	}
}

func TestGretaName(t *testing.T) {
	if ok := Greta.Name() == greta; !ok {
		t.Fatalf("%s != %s", Greta.Name(), greta)
	}
}
