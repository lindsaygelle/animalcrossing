package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRodAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Rod.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rod.Animal(), s)
	}
}

func TestRodName(t *testing.T) {
	if ok := Rod.Name() == rod; !ok {
		t.Fatalf("%s != %s", Rod.Name(), rod)
	}
}
