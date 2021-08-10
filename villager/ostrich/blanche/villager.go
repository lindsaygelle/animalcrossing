package blanche

import (
	"github.com/lindsaygelle/animalcrossing/animal/ostrich"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ost08"
)

const (
	gender string = "female"
)

const (
	id string = "blanche"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Blanche.
	Villager = villager.Villager{
		Animal:      ostrich.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Blanche,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
