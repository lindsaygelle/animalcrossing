package bulls

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestVicAnimal(t *testing.T) {
	var s string = animals.Bull.Name()
	if ok := Vic.Animal() == s; !ok {
		t.Fatalf("%s != %s", Vic.Animal(), s)
	}
}

func TestVicName(t *testing.T) {
	if ok := Vic.Name() == vic; !ok {
		t.Fatalf("%s != %s", Vic.Name(), vic)
	}
}
