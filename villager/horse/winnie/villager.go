package winnie

import (
	"github.com/lindsaygelle/animalcrossing/animal/horse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "winnie"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Winnie.
	Villager = villager.Villager{
		Animal:  horse.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
