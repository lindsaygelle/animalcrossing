package hopper

import (
	"github.com/lindsaygelle/animalcrossing/animal/penguin"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "hopper"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Hopper.
	Villager = villager.Villager{
		Animal:      penguin.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Hopper,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
