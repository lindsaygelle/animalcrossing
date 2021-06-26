package horses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBuckAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Buck.Animal() == s; !ok {
		t.Fatalf("%s != %s", Buck.Animal(), s)
	}
}

func TestBuckName(t *testing.T) {
	if ok := Buck.Name() == buck; !ok {
		t.Fatalf("%s != %s", Buck.Name(), buck)
	}
}
