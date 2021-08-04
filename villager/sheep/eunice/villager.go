package eunice

import (
	"github.com/lindsaygelle/animalcrossing/animal/sheep"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "eunice"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Eunice.
	Villager = villager.Villager{
		Animal:      sheep.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Eunice,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
