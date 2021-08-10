package ike

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "bea11"
)

const (
	gender string = "male"
)

const (
	id string = "ike"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Ike.
	Villager = villager.Villager{
		Animal:      bear.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Ike,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
