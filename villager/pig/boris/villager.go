package boris

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pig09"
)

const (
	gender string = "male"
)

const (
	id string = "boris"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Boris.
	Villager = villager.Villager{
		Animal:      pig.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Boris,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
