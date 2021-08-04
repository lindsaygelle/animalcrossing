package bud

import (
	"github.com/lindsaygelle/animalcrossing/animal/lion"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "bud"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Bud.
	Villager = villager.Villager{
		Animal:      lion.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Bud,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
