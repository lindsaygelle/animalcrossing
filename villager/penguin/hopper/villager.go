package hopper

import (
	"github.com/lindsaygelle/animalcrossing/animal/penguin"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "hopper"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Hopper.
	Villager = villager.Villager{
		Animal:  penguin.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
