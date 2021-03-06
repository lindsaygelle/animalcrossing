package teddy

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "bea00"
)

const (
	gender string = "male"
)

const (
	id string = "teddy"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Teddy.
	Villager = villager.Villager{
		Animal:      bear.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Teddy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
