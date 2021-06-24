package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBobAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Bob.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bob.Animal(), s)
	}
}

func TestBobName(t *testing.T) {
	if ok := Bob.Name() == bob; !ok {
		t.Fatalf("%s != %s", Bob.Name(), bob)
	}
}
