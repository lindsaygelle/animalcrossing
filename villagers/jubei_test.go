package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJubeiAnimal(t *testing.T) {
	var s string = animals.Lion.Name()
	if ok := Jubei.Animal() == s; !ok {
		t.Fatalf("%s != %s", Jubei.Animal(), s)
	}
}

func TestJubeiName(t *testing.T) {
	if ok := Jubei.Name() == jubei; !ok {
		t.Fatalf("%s != %s", Jubei.Name(), jubei)
	}
}
