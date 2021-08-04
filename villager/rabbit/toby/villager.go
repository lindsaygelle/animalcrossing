package toby

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "toby"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Toby.
	Villager = villager.Villager{
		Animal:      rabbit.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Toby,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
