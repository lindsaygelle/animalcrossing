package pompom

import (
	"github.com/lindsaygelle/animalcrossing/animal/duck"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "duk05"
)

const (
	gender string = "female"
)

const (
	id string = "pompom"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Pompom.
	Villager = villager.Villager{
		Animal:      duck.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Pompom,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
