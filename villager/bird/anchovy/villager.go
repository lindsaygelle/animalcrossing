package anchovy

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "brd02"
)

const (
	gender string = "male"
)

const (
	id string = "anchovy"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Anchovy.
	Villager = villager.Villager{
		Animal:      bird.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Anchovy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
