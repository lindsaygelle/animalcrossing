package bears

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJudyAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Judy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Judy.Animal(), s)
	}
}

func TestJudyName(t *testing.T) {
	if ok := Judy.Name() == judy; !ok {
		t.Fatalf("%s != %s", Judy.Name(), judy)
	}
}
