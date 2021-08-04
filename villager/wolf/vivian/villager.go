package vivian

import (
	"github.com/lindsaygelle/animalcrossing/animal/wolf"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "vivian"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Vivian.
	Villager = villager.Villager{
		Animal:      wolf.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Vivian,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
