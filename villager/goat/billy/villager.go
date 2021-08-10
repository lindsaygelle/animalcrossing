package billy

import (
	"github.com/lindsaygelle/animalcrossing/animal/goat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "goa02"
)

const (
	gender string = "male"
)

const (
	id string = "billy"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Billy.
	Villager = villager.Villager{
		Animal:      goat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Billy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
