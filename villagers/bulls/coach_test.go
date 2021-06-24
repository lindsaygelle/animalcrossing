package bulls

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCoachAnimal(t *testing.T) {
	var s string = animals.Bull.Name()
	if ok := Coach.Animal() == s; !ok {
		t.Fatalf("%s != %s", Coach.Animal(), s)
	}
}

func TestCoachName(t *testing.T) {
	if ok := Coach.Name() == coach; !ok {
		t.Fatalf("%s != %s", Coach.Name(), coach)
	}
}
