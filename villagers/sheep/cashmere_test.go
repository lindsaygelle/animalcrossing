package sheep

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCashmereAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Cashmere.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cashmere.Animal(), s)
	}
}

func TestCashmereName(t *testing.T) {
	if ok := Cashmere.Name() == cashmere; !ok {
		t.Fatalf("%s != %s", Cashmere.Name(), cashmere)
	}
}
