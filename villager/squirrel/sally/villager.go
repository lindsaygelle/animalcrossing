package sally

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "sally"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sally.
	Villager = villager.Villager{
		Animal:  squirrel.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
