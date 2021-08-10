package rasher

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pig02"
)

const (
	gender string = "male"
)

const (
	id string = "rasher"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rasher.
	Villager = villager.Villager{
		Animal:      pig.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rasher,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
