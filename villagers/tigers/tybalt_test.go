package tigers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTybaltAnimal(t *testing.T) {
	var s string = animals.Tiger.Name()
	if ok := Tybalt.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tybalt.Animal(), s)
	}
}

func TestTybaltName(t *testing.T) {
	if ok := Tybalt.Name() == tybalt; !ok {
		t.Fatalf("%s != %s", Tybalt.Name(), tybalt)
	}
}
