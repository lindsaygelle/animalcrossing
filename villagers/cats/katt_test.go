package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKattAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Katt.Animal() == s; !ok {
		t.Fatalf("%s != %s", Katt.Animal(), s)
	}
}

func TestKattName(t *testing.T) {
	if ok := Katt.Name() == katt; !ok {
		t.Fatalf("%s != %s", Katt.Name(), katt)
	}
}
