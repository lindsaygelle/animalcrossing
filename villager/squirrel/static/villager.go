package static

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "squ08"
)

const (
	gender string = "male"
)

const (
	id string = "static"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Static.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Static,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
