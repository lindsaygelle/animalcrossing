package deers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDianaAnimal(t *testing.T) {
	var s string = animals.Deer.Name()
	if ok := Diana.Animal() == s; !ok {
		t.Fatalf("%s != %s", Diana.Animal(), s)
	}
}

func TestDianaName(t *testing.T) {
	if ok := Diana.Name() == diana; !ok {
		t.Fatalf("%s != %s", Diana.Name(), diana)
	}
}
