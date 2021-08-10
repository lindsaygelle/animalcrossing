package sylvana

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "squ14"
)

const (
	gender string = "female"
)

const (
	id string = "sylvana"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sylvana.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Sylvana,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
