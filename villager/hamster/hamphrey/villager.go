package hamphrey

import (
	"github.com/lindsaygelle/animalcrossing/animal/hamster"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ham07"
)

const (
	gender string = "male"
)

const (
	id string = "hamphrey"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Hamphrey.
	Villager = villager.Villager{
		Animal:      hamster.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Hamphrey,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
