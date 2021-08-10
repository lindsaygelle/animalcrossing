package rory

import (
	"github.com/lindsaygelle/animalcrossing/animal/lion"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "lon07"
)

const (
	gender string = "male"
)

const (
	id string = "rory"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rory.
	Villager = villager.Villager{
		Animal:      lion.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rory,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
