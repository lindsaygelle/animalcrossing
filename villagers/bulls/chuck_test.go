package bulls

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChuckAnimal(t *testing.T) {
	var s string = animals.Bull.Name()
	if ok := Chuck.Animal() == s; !ok {
		t.Fatalf("%s != %s", Chuck.Animal(), s)
	}
}

func TestChuckName(t *testing.T) {
	if ok := Chuck.Name() == chuck; !ok {
		t.Fatalf("%s != %s", Chuck.Name(), chuck)
	}
}
