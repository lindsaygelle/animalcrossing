package gayle

import (
	"github.com/lindsaygelle/animalcrossing/animal/alligator"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "gayle"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Gayle.
	Villager = villager.Villager{
		Animal:      alligator.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Gayle,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
