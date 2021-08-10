package stella

import (
	"github.com/lindsaygelle/animalcrossing/animal/sheep"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "shp03"
)

const (
	gender string = "female"
)

const (
	id string = "stella"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Stella.
	Villager = villager.Villager{
		Animal:      sheep.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Stella,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
