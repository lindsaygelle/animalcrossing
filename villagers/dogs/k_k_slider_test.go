package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKKSliderAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := KKSlider.Animal() == s; !ok {
		t.Fatalf("%s != %s", KKSlider.Animal(), s)
	}
}

func TestKKSliderName(t *testing.T) {
	if ok := KKSlider.Name() == kkSlider; !ok {
		t.Fatalf("%s != %s", KKSlider.Name(), kkSlider)
	}
}
