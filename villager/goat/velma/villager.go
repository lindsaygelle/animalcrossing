package velma

import (
	"github.com/lindsaygelle/animalcrossing/animal/goat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "velma"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Velma.
	Villager = villager.Villager{
		Animal:      goat.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Velma,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
