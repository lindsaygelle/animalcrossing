package colton

import (
	"github.com/lindsaygelle/animalcrossing/animal/horse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "hrs11"
)

const (
	gender string = "male"
)

const (
	id string = "colton"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Colton.
	Villager = villager.Villager{
		Animal:      horse.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Colton,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
