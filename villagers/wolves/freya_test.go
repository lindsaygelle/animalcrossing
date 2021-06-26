package wolves

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFreyaAnimal(t *testing.T) {
	var s string = animals.Wolf.Name()
	if ok := Freya.Animal() == s; !ok {
		t.Fatalf("%s != %s", Freya.Animal(), s)
	}
}

func TestFreyaName(t *testing.T) {
	if ok := Freya.Name() == freya; !ok {
		t.Fatalf("%s != %s", Freya.Name(), freya)
	}
}
