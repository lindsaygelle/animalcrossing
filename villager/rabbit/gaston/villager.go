package gaston

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "rbt04"
)

const (
	gender string = "male"
)

const (
	id string = "gaston"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Gaston.
	Villager = villager.Villager{
		Animal:      rabbit.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Gaston,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
