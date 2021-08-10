package lily

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "flg00"
)

const (
	gender string = "female"
)

const (
	id string = "lily"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Lily.
	Villager = villager.Villager{
		Animal:      frog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Lily,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
