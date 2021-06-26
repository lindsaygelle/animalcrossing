package otters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLottieAnimal(t *testing.T) {
	var s string = animals.Otter.Name()
	if ok := Lottie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lottie.Animal(), s)
	}
}

func TestLottieName(t *testing.T) {
	if ok := Lottie.Name() == lottie; !ok {
		t.Fatalf("%s != %s", Lottie.Name(), lottie)
	}
}
