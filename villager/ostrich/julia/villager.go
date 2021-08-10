package julia

import (
	"github.com/lindsaygelle/animalcrossing/animal/ostrich"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ost05"
)

const (
	gender string = "female"
)

const (
	id string = "julia"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Julia.
	Villager = villager.Villager{
		Animal:      ostrich.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Julia,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
