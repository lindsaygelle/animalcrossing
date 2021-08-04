package quillson

import (
	"github.com/lindsaygelle/animalcrossing/animal/duck"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "quillson"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Quillson.
	Villager = villager.Villager{
		Animal:      duck.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Quillson,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
