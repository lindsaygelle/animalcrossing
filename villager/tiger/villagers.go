package tiger

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/tiger/bangle"
	"github.com/lindsaygelle/animalcrossing/villager/tiger/bianca"
	"github.com/lindsaygelle/animalcrossing/villager/tiger/claudia"
	"github.com/lindsaygelle/animalcrossing/villager/tiger/leonardo"
	"github.com/lindsaygelle/animalcrossing/villager/tiger/rolf"
	"github.com/lindsaygelle/animalcrossing/villager/tiger/rowan"
	"github.com/lindsaygelle/animalcrossing/villager/tiger/tybalt"
)

var (
	Villagers = villager.NewVillagers(
		bangle.Villager,
		bianca.Villager,
		claudia.Villager,
		leonardo.Villager,
		rolf.Villager,
		rowan.Villager,
		tybalt.Villager)
)
