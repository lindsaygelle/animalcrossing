package bigtop

import (
	"github.com/lindsaygelle/animalcrossing/animal/elephant"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = ""
)

const (
	gender string = "male"
)

const (
	id string = "bigtop"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for BigTop.
	Villager = villager.Villager{
		Animal:      elephant.Animal,
		Birthday:    birthday,
		Code:        code,
		Gender:      gender,
		Id:          id,
		Key:         villager.BigTop,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
