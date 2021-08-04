package doc

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "doc"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Doc.
	Villager = villager.Villager{
		Animal:      rabbit.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Doc,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
