package horses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWinnieAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Winnie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Winnie.Animal(), s)
	}
}

func TestWinnieName(t *testing.T) {
	if ok := Winnie.Name() == winnie; !ok {
		t.Fatalf("%s != %s", Winnie.Name(), winnie)
	}
}
