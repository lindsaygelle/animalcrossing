package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDaisyAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Daisy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Daisy.Animal(), s)
	}
}

func TestDaisyName(t *testing.T) {
	if ok := Daisy.Name() == daisy; !ok {
		t.Fatalf("%s != %s", Daisy.Name(), daisy)
	}
}
