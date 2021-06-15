package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPironkonAnimal(t *testing.T) {
	var s string = animals.Alligator.Name()
	if ok := Pironkon.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pironkon.Animal(), s)
	}
}

func TestPironkonName(t *testing.T) {
	if ok := Pironkon.Name() == pironkon; !ok {
		t.Fatalf("%s != %s", Pironkon.Name(), pironkon)
	}
}
