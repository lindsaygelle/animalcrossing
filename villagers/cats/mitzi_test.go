package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMitziAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Mitzi.Animal() == s; !ok {
		t.Fatalf("%s != %s", Mitzi.Animal(), s)
	}
}

func TestMitziName(t *testing.T) {
	if ok := Mitzi.Name() == mitzi; !ok {
		t.Fatalf("%s != %s", Mitzi.Name(), mitzi)
	}
}
