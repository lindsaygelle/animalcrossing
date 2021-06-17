package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAveryAnimal(t *testing.T) {
	var s string = animals.Eagle.Name()
	if ok := Avery.Animal() == s; !ok {
		t.Fatalf("%s != %s", Avery.Animal(), s)
	}
}

func TestAveryName(t *testing.T) {
	if ok := Avery.Name() == avery; !ok {
		t.Fatalf("%s != %s", Avery.Name(), avery)
	}
}
