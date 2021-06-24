package eagles

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFrankAnimal(t *testing.T) {
	var s string = animals.Eagle.Name()
	if ok := Frank.Animal() == s; !ok {
		t.Fatalf("%s != %s", Frank.Animal(), s)
	}
}

func TestFrankName(t *testing.T) {
	if ok := Frank.Name() == frank; !ok {
		t.Fatalf("%s != %s", Frank.Name(), frank)
	}
}
