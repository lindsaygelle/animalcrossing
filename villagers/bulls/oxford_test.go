package bulls

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestOxfordAnimal(t *testing.T) {
	var s string = animals.Bull.Name()
	if ok := Oxford.Animal() == s; !ok {
		t.Fatalf("%s != %s", Oxford.Animal(), s)
	}
}

func TestOxfordName(t *testing.T) {
	if ok := Oxford.Name() == oxford; !ok {
		t.Fatalf("%s != %s", Oxford.Name(), oxford)
	}
}
