package filbert

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "squ02"
)

const (
	gender string = "male"
)

const (
	id string = "filbert"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Filbert.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Filbert,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
