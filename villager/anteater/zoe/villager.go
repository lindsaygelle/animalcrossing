package zoe

import (
	"github.com/lindsaygelle/animalcrossing/animal/anteater"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "zoe"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Zoe.
	Villager = villager.Villager{
		Animal:      anteater.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Zoe,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
