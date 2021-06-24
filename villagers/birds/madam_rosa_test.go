package birds

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMadamRosaAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := MadamRosa.Animal() == s; !ok {
		t.Fatalf("%s != %s", MadamRosa.Animal(), s)
	}
}

func TestMadamRosaName(t *testing.T) {
	if ok := MadamRosa.Name() == madamRosa; !ok {
		t.Fatalf("%s != %s", MadamRosa.Name(), madamRosa)
	}
}