package marshal

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "squ17"
)

const (
	gender string = "male"
)

const (
	id string = "marshal"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Marshal.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Marshal,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
