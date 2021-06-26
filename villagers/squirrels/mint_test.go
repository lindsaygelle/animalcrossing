package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMintAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Mint.Animal() == s; !ok {
		t.Fatalf("%s != %s", Mint.Animal(), s)
	}
}

func TestMintName(t *testing.T) {
	if ok := Mint.Name() == mint; !ok {
		t.Fatalf("%s != %s", Mint.Name(), mint)
	}
}
