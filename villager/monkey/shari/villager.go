package shari

import (
	"github.com/lindsaygelle/animalcrossing/animal/monkey"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "mnk07"
)

const (
	gender string = "female"
)

const (
	id string = "shari"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Shari.
	Villager = villager.Villager{
		Animal:      monkey.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Shari,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
