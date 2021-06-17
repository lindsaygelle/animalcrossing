package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBonesAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Bones.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bones.Animal(), s)
	}
}

func TestBonesName(t *testing.T) {
	if ok := Bones.Name() == bones; !ok {
		t.Fatalf("%s != %s", Bones.Name(), bones)
	}
}
