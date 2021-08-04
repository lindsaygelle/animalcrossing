package aurora

import (
	"github.com/lindsaygelle/animalcrossing/animal/penguin"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "aurora"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Aurora.
	Villager = villager.Villager{
		Animal:      penguin.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Aurora,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
