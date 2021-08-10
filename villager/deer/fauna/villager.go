package fauna

import (
	"github.com/lindsaygelle/animalcrossing/animal/deer"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "der00"
)

const (
	gender string = "female"
)

const (
	id string = "fauna"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Fauna.
	Villager = villager.Villager{
		Animal:      deer.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Fauna,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
