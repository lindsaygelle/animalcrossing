package wartjr

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "wartjr"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for WartJr.
	Villager = villager.Villager{
		Animal:      frog.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.WartJr,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
