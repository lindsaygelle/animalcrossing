package ursala

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "bea08"
)

const (
	gender string = "female"
)

const (
	id string = "ursala"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Ursala.
	Villager = villager.Villager{
		Animal:      bear.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Ursala,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
