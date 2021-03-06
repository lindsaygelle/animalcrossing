package caroline

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "squ06"
)

const (
	gender string = "female"
)

const (
	id string = "caroline"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Caroline.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Caroline,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
