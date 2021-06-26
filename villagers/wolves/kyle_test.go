package wolves

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKyleAnimal(t *testing.T) {
	var s string = animals.Wolf.Name()
	if ok := Kyle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Kyle.Animal(), s)
	}
}

func TestKyleName(t *testing.T) {
	if ok := Kyle.Name() == kyle; !ok {
		t.Fatalf("%s != %s", Kyle.Name(), kyle)
	}
}
