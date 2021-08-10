package drift

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "flg04"
)

const (
	gender string = "male"
)

const (
	id string = "drift"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Drift.
	Villager = villager.Villager{
		Animal:      frog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Drift,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
