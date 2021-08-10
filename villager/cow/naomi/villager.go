package naomi

import (
	"github.com/lindsaygelle/animalcrossing/animal/cow"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cow07"
)

const (
	gender string = "female"
)

const (
	id string = "naomi"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Naomi.
	Villager = villager.Villager{
		Animal:      cow.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Naomi,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
