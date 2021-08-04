package winnie

import (
	"github.com/lindsaygelle/animalcrossing/animal/horse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "winnie"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Winnie.
	Villager = villager.Villager{
		Animal:      horse.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Winnie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
