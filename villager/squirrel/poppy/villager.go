package poppy

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "squ15"
)

const (
	gender string = "female"
)

const (
	id string = "poppy"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Poppy.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Poppy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
