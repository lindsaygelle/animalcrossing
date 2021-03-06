package muffy

import (
	"github.com/lindsaygelle/animalcrossing/animal/sheep"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "shp12"
)

const (
	gender string = "female"
)

const (
	id string = "muffy"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Muffy.
	Villager = villager.Villager{
		Animal:      sheep.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Muffy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
