package bessie

import (
	"github.com/lindsaygelle/animalcrossing/animal/cow"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "bessie"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Bessie.
	Villager = villager.Villager{
		Animal:      cow.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Bessie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
