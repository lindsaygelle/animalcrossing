package zucker

import (
	"github.com/lindsaygelle/animalcrossing/animal/octopus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ocp02"
)

const (
	gender string = "male"
)

const (
	id string = "zucker"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Zucker.
	Villager = villager.Villager{
		Animal:      octopus.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Zucker,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
