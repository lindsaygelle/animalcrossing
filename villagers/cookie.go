package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cookie string = "Cookie"
)

var (
	// Cookie is an Animal Crossing villager.
	//
	// Cookie is a Dog.
	Cookie Villager = villager{
		animal: animals.Dog.Name(),
		name:   cookie}
)
