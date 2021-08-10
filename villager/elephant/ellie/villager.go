package ellie

import (
	"github.com/lindsaygelle/animalcrossing/animal/elephant"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "elp07"
)

const (
	gender string = "female"
)

const (
	id string = "ellie"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Ellie.
	Villager = villager.Villager{
		Animal:      elephant.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Ellie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
