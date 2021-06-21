package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChrissyAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Chrissy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Chrissy.Animal(), s)
	}
}

func TestChrissyName(t *testing.T) {
	if ok := Chrissy.Name() == chrissy; !ok {
		t.Fatalf("%s != %s", Chrissy.Name(), chrissy)
	}
}
