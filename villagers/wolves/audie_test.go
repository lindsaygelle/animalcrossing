package wolves

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAudieAnimal(t *testing.T) {
	var s string = animals.Wolf.Name()
	if ok := Audie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Audie.Animal(), s)
	}
}

func TestAudieName(t *testing.T) {
	if ok := Audie.Name() == audie; !ok {
		t.Fatalf("%s != %s", Audie.Name(), audie)
	}
}
