package hedgehogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSableAnimal(t *testing.T) {
	var s string = animals.Hedgehog.Name()
	if ok := Sable.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sable.Animal(), s)
	}
}

func TestSableName(t *testing.T) {
	if ok := Sable.Name() == sable; !ok {
		t.Fatalf("%s != %s", Sable.Name(), sable)
	}
}
