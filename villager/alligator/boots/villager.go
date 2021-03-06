package boots

import (
	"github.com/lindsaygelle/animalcrossing/animal/alligator"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "crd02"
)

const (
	gender string = "male"
)

const (
	id string = "boots"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Boots.
	Villager = villager.Villager{
		Animal:      alligator.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Boots,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
