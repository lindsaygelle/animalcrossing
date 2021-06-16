package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBamAnimal(t *testing.T) {
	var s string = animals.Deer.Name()
	if ok := Bam.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bam.Animal(), s)
	}
}

func TestBamName(t *testing.T) {
	if ok := Bam.Name() == bam; !ok {
		t.Fatalf("%s != %s", Bam.Name(), bam)
	}
}
