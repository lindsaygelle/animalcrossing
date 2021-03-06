package chrissy

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "rbt13"
)

const (
	gender string = "female"
)

const (
	id string = "chrissy"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Chrissy.
	Villager = villager.Villager{
		Animal:      rabbit.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Chrissy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
