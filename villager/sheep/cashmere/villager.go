package cashmere

import (
	"github.com/lindsaygelle/animalcrossing/animal/sheep"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "shp04"
)

const (
	gender string = "female"
)

const (
	id string = "cashmere"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cashmere.
	Villager = villager.Villager{
		Animal:      sheep.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Cashmere,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
