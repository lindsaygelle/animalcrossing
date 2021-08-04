package emerald

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "emerald"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Emerald.
	Villager = villager.Villager{
		Animal:      frog.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Emerald,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
