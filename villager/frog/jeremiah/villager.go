package jeremiah

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "flg07"
)

const (
	gender string = "male"
)

const (
	id string = "jeremiah"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Jeremiah.
	Villager = villager.Villager{
		Animal:      frog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Jeremiah,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
