package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCheriAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Cheri.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cheri.Animal(), s)
	}
}

func TestCheriName(t *testing.T) {
	if ok := Cheri.Name() == cheri; !ok {
		t.Fatalf("%s != %s", Cheri.Name(), cheri)
	}
}
