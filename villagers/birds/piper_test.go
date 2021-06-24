package birds

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPiperAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Piper.Animal() == s; !ok {
		t.Fatalf("%s != %s", Piper.Animal(), s)
	}
}

func TestPiperName(t *testing.T) {
	if ok := Piper.Name() == piper; !ok {
		t.Fatalf("%s != %s", Piper.Name(), piper)
	}
}
