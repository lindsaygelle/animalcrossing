package snake

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "rbt03"
)

const (
	gender string = "male"
)

const (
	id string = "snake"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Snake.
	Villager = villager.Villager{
		Animal:      rabbit.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Snake,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
