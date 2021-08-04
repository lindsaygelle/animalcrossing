package marina

import (
	"github.com/lindsaygelle/animalcrossing/animal/octopus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "marina"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Marina.
	Villager = villager.Villager{
		Animal:      octopus.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Marina,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
