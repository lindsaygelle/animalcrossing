package stitches

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cbr05"
)

const (
	gender string = "male"
)

const (
	id string = "stitches"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Stitches.
	Villager = villager.Villager{
		Animal:      bearcub.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Stitches,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
