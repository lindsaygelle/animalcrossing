package molly

import (
	"github.com/lindsaygelle/animalcrossing/animal/duck"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "molly"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Molly.
	Villager = villager.Villager{
		Animal:      duck.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Molly,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
