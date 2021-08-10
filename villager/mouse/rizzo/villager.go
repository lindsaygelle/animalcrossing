package rizzo

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "mus09"
)

const (
	gender string = "male"
)

const (
	id string = "rizzo"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rizzo.
	Villager = villager.Villager{
		Animal:      mouse.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rizzo,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
