package kody

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "kody"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Kody.
	Villager = villager.Villager{
		Animal:  bearcub.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
