package cube

import (
	"github.com/lindsaygelle/animalcrossing/animal/penguin"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pgn02"
)

const (
	gender string = "male"
)

const (
	id string = "cube"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cube.
	Villager = villager.Villager{
		Animal:      penguin.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Cube,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
