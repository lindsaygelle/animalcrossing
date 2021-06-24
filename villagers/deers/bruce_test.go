package deers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBruceAnimal(t *testing.T) {
	var s string = animals.Deer.Name()
	if ok := Bruce.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bruce.Animal(), s)
	}
}

func TestBruceName(t *testing.T) {
	if ok := Bruce.Name() == bruce; !ok {
		t.Fatalf("%s != %s", Bruce.Name(), bruce)
	}
}
