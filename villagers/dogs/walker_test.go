package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWalkerAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Walker.Animal() == s; !ok {
		t.Fatalf("%s != %s", Walker.Animal(), s)
	}
}

func TestWalkerName(t *testing.T) {
	if ok := Walker.Name() == walker; !ok {
		t.Fatalf("%s != %s", Walker.Name(), walker)
	}
}
