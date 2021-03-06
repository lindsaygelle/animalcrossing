package dobie

import (
	"github.com/lindsaygelle/animalcrossing/animal/wolf"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "wol04"
)

const (
	gender string = "male"
)

const (
	id string = "dobie"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Dobie.
	Villager = villager.Villager{
		Animal:      wolf.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Dobie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
