package canberra

import (
	"github.com/lindsaygelle/animalcrossing/animal/koala"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "kal08"
)

const (
	gender string = "female"
)

const (
	id string = "canberra"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Canberra.
	Villager = villager.Villager{
		Animal:      koala.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Canberra,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
