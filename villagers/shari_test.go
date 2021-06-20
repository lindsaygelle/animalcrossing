package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestShariAnimal(t *testing.T) {
	var s string = animals.Monkey.Name()
	if ok := Shari.Animal() == s; !ok {
		t.Fatalf("%s != %s", Shari.Animal(), s)
	}
}

func TestShariName(t *testing.T) {
	if ok := Shari.Name() == shari; !ok {
		t.Fatalf("%s != %s", Shari.Name(), shari)
	}
}
