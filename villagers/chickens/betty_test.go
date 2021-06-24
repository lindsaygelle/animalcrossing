package chickens

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBettyAnimal(t *testing.T) {
	var s string = animals.Chicken.Name()
	if ok := Betty.Animal() == s; !ok {
		t.Fatalf("%s != %s", Betty.Animal(), s)
	}
}

func TestBettyName(t *testing.T) {
	if ok := Betty.Name() == betty; !ok {
		t.Fatalf("%s != %s", Betty.Name(), betty)
	}
}
