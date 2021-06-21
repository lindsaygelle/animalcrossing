package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDomAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Dom.Animal() == s; !ok {
		t.Fatalf("%s != %s", Dom.Animal(), s)
	}
}

func TestDomName(t *testing.T) {
	if ok := Dom.Name() == dom; !ok {
		t.Fatalf("%s != %s", Dom.Name(), dom)
	}
}
