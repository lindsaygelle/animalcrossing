package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBenjaminAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Benjamin.Animal() == s; !ok {
		t.Fatalf("%s != %s", Benjamin.Animal(), s)
	}
}

func TestBenjaminName(t *testing.T) {
	if ok := Benjamin.Name() == benjamin; !ok {
		t.Fatalf("%s != %s", Benjamin.Name(), benjamin)
	}
}
