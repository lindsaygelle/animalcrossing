package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKittyAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Kitty.Animal() == s; !ok {
		t.Fatalf("%s != %s", Kitty.Animal(), s)
	}
}

func TestKittyName(t *testing.T) {
	if ok := Kitty.Name() == kitty; !ok {
		t.Fatalf("%s != %s", Kitty.Name(), kitty)
	}
}
