package giraffes

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGracieAnimal(t *testing.T) {
	var s string = animals.Giraffe.Name()
	if ok := Gracie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gracie.Animal(), s)
	}
}

func TestGracieName(t *testing.T) {
	if ok := Gracie.Name() == gracie; !ok {
		t.Fatalf("%s != %s", Gracie.Name(), gracie)
	}
}
