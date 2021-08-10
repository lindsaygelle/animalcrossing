package peggy

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pig11"
)

const (
	gender string = "female"
)

const (
	id string = "peggy"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Peggy.
	Villager = villager.Villager{
		Animal:      pig.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Peggy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
