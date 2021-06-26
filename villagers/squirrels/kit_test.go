package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKitAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Kit.Animal() == s; !ok {
		t.Fatalf("%s != %s", Kit.Animal(), s)
	}
}

func TestKitName(t *testing.T) {
	if ok := Kit.Name() == kit; !ok {
		t.Fatalf("%s != %s", Kit.Name(), kit)
	}
}
