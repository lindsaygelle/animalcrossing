package birds

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAdmiralAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Admiral.Animal() == s; !ok {
		t.Fatalf("%s != %s", Admiral.Animal(), s)
	}
}

func TestAdmiralName(t *testing.T) {
	if ok := Admiral.Name() == admiral; !ok {
		t.Fatalf("%s != %s", Admiral.Name(), admiral)
	}
}
