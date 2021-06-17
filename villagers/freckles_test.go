package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFrecklesAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Freckles.Animal() == s; !ok {
		t.Fatalf("%s != %s", Freckles.Animal(), s)
	}
}

func TestFrecklesName(t *testing.T) {
	if ok := Freckles.Name() == freckles; !ok {
		t.Fatalf("%s != %s", Freckles.Name(), freckles)
	}
}
