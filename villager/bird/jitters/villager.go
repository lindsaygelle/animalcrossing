package jitters

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "jitters"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Jitters.
	Villager = villager.Villager{
		Animal:  bird.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
