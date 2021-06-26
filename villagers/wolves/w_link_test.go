package wolves

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWLinkAnimal(t *testing.T) {
	var s string = animals.Wolf.Name()
	if ok := WLink.Animal() == s; !ok {
		t.Fatalf("%s != %s", WLink.Animal(), s)
	}
}

func TestWLinkName(t *testing.T) {
	if ok := WLink.Name() == wLink; !ok {
		t.Fatalf("%s != %s", WLink.Name(), wLink)
	}
}
