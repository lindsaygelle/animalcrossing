package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPudgeName(t *testing.T) {
	if ok := Pudge.Name() == pudge; !ok {
		t.Fatalf("%s != %s", Pudge.Name(), pudge)
	}
}

func TestPudgeSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Pudge.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pudge.Animal(), s)
	}
}
