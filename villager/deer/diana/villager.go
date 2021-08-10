package diana

import (
	"github.com/lindsaygelle/animalcrossing/animal/deer"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "der08"
)

const (
	gender string = "female"
)

const (
	id string = "diana"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Diana.
	Villager = villager.Villager{
		Animal:      deer.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Diana,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
