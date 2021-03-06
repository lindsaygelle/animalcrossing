package flossie

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = ""
)

const (
	gender string = "female"
)

const (
	id string = "flossie"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Flossie.
	Villager = villager.Villager{
		Animal:      mouse.Animal,
		Birthday:    birthday,
		Code:        code,
		Gender:      gender,
		Id:          id,
		Key:         villager.Flossie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
