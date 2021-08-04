package carrot

import (
	"github.com/lindsaygelle/animalcrossing/animal/cow"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "carrot"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Carrot.
	Villager = villager.Villager{
		Animal:      cow.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Carrot,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
