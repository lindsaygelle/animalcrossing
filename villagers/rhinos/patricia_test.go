package rhinoceroses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPatriciaAnimal(t *testing.T) {
	var s string = animals.Rhinoceros.Name()
	if ok := Patricia.Animal() == s; !ok {
		t.Fatalf("%s != %s", Patricia.Animal(), s)
	}
}

func TestPatriciaName(t *testing.T) {
	if ok := Patricia.Name() == patricia; !ok {
		t.Fatalf("%s != %s", Patricia.Name(), patricia)
	}
}
