package mira

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "rbt19"
)

const (
	gender string = "female"
)

const (
	id string = "mira"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Mira.
	Villager = villager.Villager{
		Animal:      rabbit.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Mira,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
