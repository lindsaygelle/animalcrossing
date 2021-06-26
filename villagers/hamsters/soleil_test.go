package hamsters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSoleilAnimal(t *testing.T) {
	var s string = animals.Hamster.Name()
	if ok := Soleil.Animal() == s; !ok {
		t.Fatalf("%s != %s", Soleil.Animal(), s)
	}
}

func TestSoleilName(t *testing.T) {
	if ok := Soleil.Name() == soleil; !ok {
		t.Fatalf("%s != %s", Soleil.Name(), soleil)
	}
}
