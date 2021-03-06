package admiral

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "brd06"
)

const (
	gender string = "male"
)

const (
	id string = "admiral"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Admiral.
	Villager = villager.Villager{
		Animal:      bird.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Admiral,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
