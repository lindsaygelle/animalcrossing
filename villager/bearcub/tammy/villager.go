package tammy

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cbr17"
)

const (
	gender string = "female"
)

const (
	id string = "tammy"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tammy.
	Villager = villager.Villager{
		Animal:      bearcub.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Tammy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
