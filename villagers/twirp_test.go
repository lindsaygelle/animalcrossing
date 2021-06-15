package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTwirpAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Twirp.Animal() == s; !ok {
		t.Fatalf("%s != %s", Twirp.Animal(), s)

	}
}

func TestTwirpName(t *testing.T) {
	if ok := Twirp.Name() == twirp; !ok {
		t.Fatalf("%s != %s", Twirp.Name(), twirp)
	}
}
