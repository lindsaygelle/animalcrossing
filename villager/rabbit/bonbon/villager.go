package bonbon

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "rbt17"
)

const (
	gender string = "female"
)

const (
	id string = "bonbon"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Bonbon.
	Villager = villager.Villager{
		Animal:      rabbit.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Bonbon,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
