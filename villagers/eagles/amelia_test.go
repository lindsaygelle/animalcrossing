package eagles

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAmeliaAnimal(t *testing.T) {
	var s string = animals.Eagle.Name()
	if ok := Amelia.Animal() == s; !ok {
		t.Fatalf("%s != %s", Amelia.Animal(), s)
	}
}

func TestAmeliaName(t *testing.T) {
	if ok := Amelia.Name() == amelia; !ok {
		t.Fatalf("%s != %s", Amelia.Name(), amelia)
	}
}
