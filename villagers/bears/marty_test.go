package bears

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMartyAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Marty.Animal() == s; !ok {
		t.Fatalf("%s != %s", Marty.Animal(), s)
	}
}

func TestMartyName(t *testing.T) {
	if ok := Marty.Name() == marty; !ok {
		t.Fatalf("%s != %s", Marty.Name(), marty)
	}
}
