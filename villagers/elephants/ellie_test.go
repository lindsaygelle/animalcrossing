package elephants

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestEllieAnimal(t *testing.T) {
	var s string = animals.Elephant.Name()
	if ok := Ellie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ellie.Animal(), s)
	}
}

func TestEllieName(t *testing.T) {
	if ok := Ellie.Name() == ellie; !ok {
		t.Fatalf("%s != %s", Ellie.Name(), ellie)
	}
}
