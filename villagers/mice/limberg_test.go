package mice

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLimbergAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Limberg.Animal() == s; !ok {
		t.Fatalf("%s != %s", Limberg.Animal(), s)
	}
}

func TestLimbergName(t *testing.T) {
	if ok := Limberg.Name() == limberg; !ok {
		t.Fatalf("%s != %s", Limberg.Name(), limberg)
	}
}
