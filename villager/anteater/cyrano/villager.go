package cyrano

import (
	"github.com/lindsaygelle/animalcrossing/animal/anteater"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "cyrano"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cyrano.
	Villager = villager.Villager{
		Animal:      anteater.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Cyrano,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
