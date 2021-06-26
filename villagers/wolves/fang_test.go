package wolves

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFangAnimal(t *testing.T) {
	var s string = animals.Wolf.Name()
	if ok := Fang.Animal() == s; !ok {
		t.Fatalf("%s != %s", Fang.Animal(), s)
	}
}

func TestFangName(t *testing.T) {
	if ok := Fang.Name() == fang; !ok {
		t.Fatalf("%s != %s", Fang.Name(), fang)
	}
}
