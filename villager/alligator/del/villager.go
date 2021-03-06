package del

import (
	"github.com/lindsaygelle/animalcrossing/animal/alligator"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "crd04"
)

const (
	gender string = "male"
)

const (
	id string = "del"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Del.
	Villager = villager.Villager{
		Animal:      alligator.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Del,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
