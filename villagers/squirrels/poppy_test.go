package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPoppyAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Poppy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Poppy.Animal(), s)
	}
}

func TestPoppyName(t *testing.T) {
	if ok := Poppy.Name() == poppy; !ok {
		t.Fatalf("%s != %s", Poppy.Name(), poppy)
	}
}
