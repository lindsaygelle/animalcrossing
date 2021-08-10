package raddle

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "flg15"
)

const (
	gender string = "male"
)

const (
	id string = "raddle"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Raddle.
	Villager = villager.Villager{
		Animal:      frog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Raddle,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
