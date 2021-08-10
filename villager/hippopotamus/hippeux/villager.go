package hippeux

import (
	"github.com/lindsaygelle/animalcrossing/animal/hippopotamus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "hip09"
)

const (
	gender string = "male"
)

const (
	id string = "hippeux"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Hippeux.
	Villager = villager.Villager{
		Animal:      hippopotamus.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Hippeux,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
