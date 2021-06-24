package birds

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCarloAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Carlo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Carlo.Animal(), s)
	}
}

func TestCarloName(t *testing.T) {
	if ok := Carlo.Name() == carlo; !ok {
		t.Fatalf("%s != %s", Carlo.Name(), carlo)
	}
}
