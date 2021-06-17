package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWartJRAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := WartJR.Animal() == s; !ok {
		t.Fatalf("%s != %s", WartJR.Animal(), s)
	}
}

func TestWartJRName(t *testing.T) {
	if ok := WartJR.Name() == wartJr; !ok {
		t.Fatalf("%s != %s", WartJR.Name(), wartJr)
	}
}
