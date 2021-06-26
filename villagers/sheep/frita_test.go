package sheep

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFritaAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Frita.Animal() == s; !ok {
		t.Fatalf("%s != %s", Frita.Animal(), s)
	}
}

func TestFritaName(t *testing.T) {
	if ok := Frita.Name() == frita; !ok {
		t.Fatalf("%s != %s", Frita.Name(), frita)
	}
}
