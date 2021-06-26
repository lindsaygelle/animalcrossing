package koalas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCanberraAnimal(t *testing.T) {
	var s string = animals.Koala.Name()
	if ok := Canberra.Animal() == s; !ok {
		t.Fatalf("%s != %s", Canberra.Animal(), s)
	}
}

func TestCanberraName(t *testing.T) {
	if ok := Canberra.Name() == canberra; !ok {
		t.Fatalf("%s != %s", Canberra.Name(), canberra)
	}
}
