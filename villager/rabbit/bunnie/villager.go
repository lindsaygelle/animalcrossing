package bunnie

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "rbt00"
)

const (
	gender string = "female"
)

const (
	id string = "bunnie"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Bunnie.
	Villager = villager.Villager{
		Animal:      rabbit.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Bunnie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
