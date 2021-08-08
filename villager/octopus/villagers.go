package octopus

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/octopus/inkwell"
	"github.com/lindsaygelle/animalcrossing/villager/octopus/marina"
	"github.com/lindsaygelle/animalcrossing/villager/octopus/octavian"
	"github.com/lindsaygelle/animalcrossing/villager/octopus/zucker"
)

var (
	Villagers = villager.NewVillagers(
		inkwell.Villager,
		marina.Villager,
		octavian.Villager,
		zucker.Villager)
)
