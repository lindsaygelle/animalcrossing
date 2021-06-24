package frogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRibbotAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Ribbot.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ribbot.Animal(), s)
	}
}

func TestRibbotName(t *testing.T) {
	if ok := Ribbot.Name() == ribbot; !ok {
		t.Fatalf("%s != %s", Ribbot.Name(), ribbot)
	}
}
