package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestZipperTBunnyAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := ZipperTBunny.Animal() == s; !ok {
		t.Fatalf("%s != %s", ZipperTBunny.Animal(), s)
	}
}

func TestZipperTBunnyName(t *testing.T) {
	if ok := ZipperTBunny.Name() == zipperTBunny; !ok {
		t.Fatalf("%s != %s", ZipperTBunny.Name(), zipperTBunny)
	}
}
