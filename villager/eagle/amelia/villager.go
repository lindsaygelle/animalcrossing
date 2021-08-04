package amelia

import (
	"github.com/lindsaygelle/animalcrossing/animal/eagle"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "amelia"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Amelia.
	Villager = villager.Villager{
		Animal:      eagle.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Amelia,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
