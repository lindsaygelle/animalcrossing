package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMarshalAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Marshal.Animal() == s; !ok {
		t.Fatalf("%s != %s", Marshal.Animal(), s)
	}
}

func TestMarshalName(t *testing.T) {
	if ok := Marshal.Name() == marshal; !ok {
		t.Fatalf("%s != %s", Marshal.Name(), marshal)
	}
}
