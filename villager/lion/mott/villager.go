package mott

import (
	"github.com/lindsaygelle/animalcrossing/animal/lion"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "lon06"
)

const (
	gender string = "male"
)

const (
	id string = "mott"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Mott.
	Villager = villager.Villager{
		Animal:      lion.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Mott,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
