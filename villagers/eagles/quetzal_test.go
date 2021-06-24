package eagles

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestQuetzalAnimal(t *testing.T) {
	var s string = animals.Eagle.Name()
	if ok := Quetzal.Animal() == s; !ok {
		t.Fatalf("%s != %s", Quetzal.Animal(), s)
	}
}

func TestQuetzalName(t *testing.T) {
	if ok := Quetzal.Name() == quetzal; !ok {
		t.Fatalf("%s != %s", Quetzal.Name(), quetzal)
	}
}
