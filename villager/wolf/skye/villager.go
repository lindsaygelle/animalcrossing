package skye

import (
	"github.com/lindsaygelle/animalcrossing/animal/wolf"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "skye"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Skye.
	Villager = villager.Villager{
		Animal:      wolf.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Skye,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
