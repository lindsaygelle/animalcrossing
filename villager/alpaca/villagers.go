package alpaca

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/alpaca/cyrus"
	"github.com/lindsaygelle/animalcrossing/villager/alpaca/reese"
)

var (
	Villagers = villager.NewVillagers(
		cyrus.Villager,
		reese.Villager)
)
