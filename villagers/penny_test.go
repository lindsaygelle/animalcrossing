package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPennyAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Penny.Animal() == s; !ok {
		t.Fatalf("%s != %s", Penny.Animal(), s)
	}
}

func TestPennyName(t *testing.T) {
	if ok := Penny.Name() == penny; !ok {
		t.Fatalf("%s != %s", Penny.Name(), penny)
	}
}
