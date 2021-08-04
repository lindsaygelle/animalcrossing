package yodel

import (
	"github.com/lindsaygelle/animalcrossing/animal/gorilla"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "yodel"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Yodel.
	Villager = villager.Villager{
		Animal:      gorilla.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Yodel,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
