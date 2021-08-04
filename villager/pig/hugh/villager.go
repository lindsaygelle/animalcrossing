package hugh

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "hugh"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Hugh.
	Villager = villager.Villager{
		Animal:      pig.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Hugh,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
