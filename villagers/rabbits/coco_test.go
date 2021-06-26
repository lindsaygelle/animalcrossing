package rabbits

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCocoAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Coco.Animal() == s; !ok {
		t.Fatalf("%s != %s", Coco.Animal(), s)
	}
}

func TestCocoName(t *testing.T) {
	if ok := Coco.Name() == coco; !ok {
		t.Fatalf("%s != %s", Coco.Name(), coco)
	}
}
