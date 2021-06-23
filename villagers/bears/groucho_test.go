package bears

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGrouchoAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Groucho.Animal() == s; !ok {
		t.Fatalf("%s != %s", Groucho.Animal(), s)
	}
}

func TestGrouchoName(t *testing.T) {
	if ok := Groucho.Name() == groucho; !ok {
		t.Fatalf("%s != %s", Groucho.Name(), groucho)
	}
}
