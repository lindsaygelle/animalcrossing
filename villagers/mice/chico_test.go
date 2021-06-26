package mice

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChicoAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Chico.Animal() == s; !ok {
		t.Fatalf("%s != %s", Chico.Animal(), s)
	}
}

func TestChicoName(t *testing.T) {
	if ok := Chico.Name() == chico; !ok {
		t.Fatalf("%s != %s", Chico.Name(), chico)
	}
}
