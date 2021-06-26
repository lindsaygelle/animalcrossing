package pelicans

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPhyllisAnimal(t *testing.T) {
	var s string = animals.Pelican.Name()
	if ok := Phyllis.Animal() == s; !ok {
		t.Fatalf("%s != %s", Phyllis.Animal(), s)
	}
}

func TestPhyllisName(t *testing.T) {
	if ok := Phyllis.Name() == phyllis; !ok {
		t.Fatalf("%s != %s", Phyllis.Name(), phyllis)
	}
}
