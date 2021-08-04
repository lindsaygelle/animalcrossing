package norma

import (
	"github.com/lindsaygelle/animalcrossing/animal/cow"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "norma"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Norma.
	Villager = villager.Villager{
		Animal:      cow.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Norma,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
