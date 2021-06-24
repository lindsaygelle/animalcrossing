package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHarveyAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Harvey.Animal() == s; !ok {
		t.Fatalf("%s != %s", Harvey.Animal(), s)
	}
}

func TestHarveyName(t *testing.T) {
	if ok := Harvey.Name() == harvey; !ok {
		t.Fatalf("%s != %s", Harvey.Name(), harvey)
	}
}
