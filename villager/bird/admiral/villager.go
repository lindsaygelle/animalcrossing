package admiral

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "admiral"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Admiral.
	Villager = villager.Villager{
		Animal:  bird.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
