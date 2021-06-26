package horses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPeachesAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Peaches.Animal() == s; !ok {
		t.Fatalf("%s != %s", Peaches.Animal(), s)
	}
}

func TestPeachesName(t *testing.T) {
	if ok := Peaches.Name() == peaches; !ok {
		t.Fatalf("%s != %s", Peaches.Name(), peaches)
	}
}
