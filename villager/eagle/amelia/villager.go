package amelia

import (
	"github.com/lindsaygelle/animalcrossing/animal/eagle"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pbr01"
)

const (
	gender string = "female"
)

const (
	id string = "amelia"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Amelia.
	Villager = villager.Villager{
		Animal:      eagle.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Amelia,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
