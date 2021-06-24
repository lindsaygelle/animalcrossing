package chameleons

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFlickAnimal(t *testing.T) {
	var s string = animals.Chameleon.Name()
	if ok := Flick.Animal() == s; !ok {
		t.Fatalf("%s != %s", Flick.Animal(), s)
	}
}

func TestFlickName(t *testing.T) {
	if ok := Flick.Name() == flick; !ok {
		t.Fatalf("%s != %s", Flick.Name(), flick)
	}
}
