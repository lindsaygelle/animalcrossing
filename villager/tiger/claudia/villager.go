package claudia

import (
	"github.com/lindsaygelle/animalcrossing/animal/tiger"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "tig05"
)

const (
	gender string = "female"
)

const (
	id string = "claudia"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Claudia.
	Villager = villager.Villager{
		Animal:      tiger.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Claudia,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
