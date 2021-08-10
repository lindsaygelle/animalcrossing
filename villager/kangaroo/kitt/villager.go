package kitt

import (
	"github.com/lindsaygelle/animalcrossing/animal/kangaroo"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "kgr00"
)

const (
	gender string = "female"
)

const (
	id string = "kitt"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Kitt.
	Villager = villager.Villager{
		Animal:      kangaroo.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Kitt,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
