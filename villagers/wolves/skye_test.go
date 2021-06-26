package wolves

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSkyeAnimal(t *testing.T) {
	var s string = animals.Wolf.Name()
	if ok := Skye.Animal() == s; !ok {
		t.Fatalf("%s != %s", Skye.Animal(), s)
	}
}

func TestSkyeName(t *testing.T) {
	if ok := Skye.Name() == skye; !ok {
		t.Fatalf("%s != %s", Skye.Name(), skye)
	}
}
