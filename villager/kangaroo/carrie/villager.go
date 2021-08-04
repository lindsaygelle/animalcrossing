package carrie

import (
	"github.com/lindsaygelle/animalcrossing/animal/kangaroo"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "carrie"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Carrie.
	Villager = villager.Villager{
		Animal:      kangaroo.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Carrie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
