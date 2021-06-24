package elephants

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTuckerAnimal(t *testing.T) {
	var s string = animals.Elephant.Name()
	if ok := Tucker.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tucker.Animal(), s)
	}
}

func TestTuckerName(t *testing.T) {
	if ok := Tucker.Name() == tucker; !ok {
		t.Fatalf("%s != %s", Tucker.Name(), tucker)
	}
}
