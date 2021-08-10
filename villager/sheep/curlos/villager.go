package curlos

import (
	"github.com/lindsaygelle/animalcrossing/animal/sheep"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "shp08"
)

const (
	gender string = "male"
)

const (
	id string = "curlos"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Curlos.
	Villager = villager.Villager{
		Animal:      sheep.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Curlos,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
