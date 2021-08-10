package pippy

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "rbt06"
)

const (
	gender string = "female"
)

const (
	id string = "pippy"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Pippy.
	Villager = villager.Villager{
		Animal:      rabbit.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Pippy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
