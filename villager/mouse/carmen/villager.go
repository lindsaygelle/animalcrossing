package carmen

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "rbt16"
)

const (
	gender string = "female"
)

const (
	id string = "carmen"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Carmen.
	Villager = villager.Villager{
		Animal:      mouse.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Carmen,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
