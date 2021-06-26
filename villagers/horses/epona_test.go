package horses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestEponaAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Epona.Animal() == s; !ok {
		t.Fatalf("%s != %s", Epona.Animal(), s)
	}
}

func TestEponaName(t *testing.T) {
	if ok := Epona.Name() == epona; !ok {
		t.Fatalf("%s != %s", Epona.Name(), epona)
	}
}
