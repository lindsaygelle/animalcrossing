package tucker

import (
	"github.com/lindsaygelle/animalcrossing/animal/elephant"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "elp09"
)

const (
	gender string = "male"
)

const (
	id string = "tucker"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tucker.
	Villager = villager.Villager{
		Animal:      elephant.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Tucker,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
