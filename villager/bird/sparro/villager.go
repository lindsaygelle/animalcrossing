package sparro

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "brd18"
)

const (
	gender string = "male"
)

const (
	id string = "sparro"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sparro.
	Villager = villager.Villager{
		Animal:      bird.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Sparro,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
