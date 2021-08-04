package charlise

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "charlise"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Charlise.
	Villager = villager.Villager{
		Animal:      bear.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Charlise,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
