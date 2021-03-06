package egbert

import (
	"github.com/lindsaygelle/animalcrossing/animal/chicken"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "chn02"
)

const (
	gender string = "male"
)

const (
	id string = "egbert"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Egbert.
	Villager = villager.Villager{
		Animal:      chicken.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Egbert,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
