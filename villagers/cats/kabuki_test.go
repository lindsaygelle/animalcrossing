package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKabukiAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Kabuki.Animal() == s; !ok {
		t.Fatalf("%s != %s", Kabuki.Animal(), s)
	}
}

func TestKabukiName(t *testing.T) {
	if ok := Kabuki.Name() == kabuki; !ok {
		t.Fatalf("%s != %s", Kabuki.Name(), kabuki)
	}
}
