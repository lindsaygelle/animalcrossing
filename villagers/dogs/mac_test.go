package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMacAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Mac.Animal() == s; !ok {
		t.Fatalf("%s != %s", Mac.Animal(), s)
	}
}

func TestMacName(t *testing.T) {
	if ok := Mac.Name() == mac; !ok {
		t.Fatalf("%s != %s", Mac.Name(), mac)
	}
}
