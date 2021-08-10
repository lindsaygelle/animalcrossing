package wendy

import (
	"github.com/lindsaygelle/animalcrossing/animal/sheep"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "shp09"
)

const (
	gender string = "female"
)

const (
	id string = "wendy"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Wendy.
	Villager = villager.Villager{
		Animal:      sheep.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Wendy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
