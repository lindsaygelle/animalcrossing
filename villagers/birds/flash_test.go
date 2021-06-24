package birds

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFlashAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Flash.Animal() == s; !ok {
		t.Fatalf("%s != %s", Flash.Animal(), s)
	}
}

func TestFlashName(t *testing.T) {
	if ok := Flash.Name() == flash; !ok {
		t.Fatalf("%s != %s", Flash.Name(), flash)
	}
}
