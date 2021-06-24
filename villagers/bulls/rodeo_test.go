package bulls

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRodeoAnimal(t *testing.T) {
	var s string = animals.Bull.Name()
	if ok := Rodeo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rodeo.Animal(), s)
	}
}

func TestRodeoName(t *testing.T) {
	if ok := Rodeo.Name() == rodeo; !ok {
		t.Fatalf("%s != %s", Rodeo.Name(), rodeo)
	}
}
