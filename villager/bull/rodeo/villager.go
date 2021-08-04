package rodeo

import (
	"github.com/lindsaygelle/animalcrossing/animal/bull"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "rodeo"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rodeo.
	Villager = villager.Villager{
		Animal:      bull.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rodeo,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
