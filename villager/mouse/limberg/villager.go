package limberg

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "mus01"
)

const (
	gender string = "male"
)

const (
	id string = "limberg"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Limberg.
	Villager = villager.Villager{
		Animal:      mouse.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Limberg,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
