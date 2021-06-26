package sea_gulls

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGullivarrrAnimal(t *testing.T) {
	var s string = animals.SeaGull.Name()
	if ok := Gullivarrr.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gullivarrr.Animal(), s)
	}
}

func TestGullivarrrName(t *testing.T) {
	if ok := Gullivarrr.Name() == gullivarrr; !ok {
		t.Fatalf("%s != %s", Gullivarrr.Name(), gullivarrr)
	}
}
