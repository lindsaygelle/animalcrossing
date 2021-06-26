package hippopotamuses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBerthaAnimal(t *testing.T) {
	var s string = animals.Hippopotamus.Name()
	if ok := Bertha.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bertha.Animal(), s)
	}
}

func TestBerthaName(t *testing.T) {
	if ok := Bertha.Name() == bertha; !ok {
		t.Fatalf("%s != %s", Bertha.Name(), bertha)
	}
}
