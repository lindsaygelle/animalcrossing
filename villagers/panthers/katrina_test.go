package panthers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKatrinaAnimal(t *testing.T) {
	var s string = animals.Panther.Name()
	if ok := Katrina.Animal() == s; !ok {
		t.Fatalf("%s != %s", Katrina.Animal(), s)
	}
}

func TestKatrinaName(t *testing.T) {
	if ok := Katrina.Name() == katrina; !ok {
		t.Fatalf("%s != %s", Katrina.Name(), katrina)
	}
}
