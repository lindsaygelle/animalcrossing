package olivia

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "olivia"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Olivia.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Olivia,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
