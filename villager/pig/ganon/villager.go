package ganon

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pig18"
)

const (
	gender string = "male"
)

const (
	id string = "ganon"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Ganon.
	Villager = villager.Villager{
		Animal:      pig.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Ganon,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
