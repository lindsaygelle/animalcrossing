package jitters

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "brd04"
)

const (
	gender string = "male"
)

const (
	id string = "jitters"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Jitters.
	Villager = villager.Villager{
		Animal:      bird.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Jitters,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
