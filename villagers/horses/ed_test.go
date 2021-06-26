package horses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestEdAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Ed.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ed.Animal(), s)
	}
}

func TestEdName(t *testing.T) {
	if ok := Ed.Name() == ed; !ok {
		t.Fatalf("%s != %s", Ed.Name(), ed)
	}
}
