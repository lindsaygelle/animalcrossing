package mint

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "mint"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Mint.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Mint,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
