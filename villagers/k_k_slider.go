package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kkSlider string = "K.K. Slider"
)

var (
	// K.K. Slider is an Animal Crossing villager.
	//
	// K.K. Slider is a Dog.
	KKSlider Villager = villager{
		animal: animals.Dog.Name(),
		name:   kkSlider}
)
