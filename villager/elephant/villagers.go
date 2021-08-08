package elephant

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/elephant/axel"
	"github.com/lindsaygelle/animalcrossing/villager/elephant/bigtop"
	"github.com/lindsaygelle/animalcrossing/villager/elephant/chai"
	"github.com/lindsaygelle/animalcrossing/villager/elephant/cyd"
	"github.com/lindsaygelle/animalcrossing/villager/elephant/dizzy"
	"github.com/lindsaygelle/animalcrossing/villager/elephant/elina"
	"github.com/lindsaygelle/animalcrossing/villager/elephant/ellie"
	"github.com/lindsaygelle/animalcrossing/villager/elephant/eloise"
	"github.com/lindsaygelle/animalcrossing/villager/elephant/margie"
	"github.com/lindsaygelle/animalcrossing/villager/elephant/opal"
	"github.com/lindsaygelle/animalcrossing/villager/elephant/paolo"
	"github.com/lindsaygelle/animalcrossing/villager/elephant/tia"
	"github.com/lindsaygelle/animalcrossing/villager/elephant/tucker"
)

var (
	Villagers = villager.NewVillagers(
		axel.Villager,
		bigtop.Villager,
		chai.Villager,
		cyd.Villager,
		dizzy.Villager,
		elina.Villager,
		ellie.Villager,
		eloise.Villager,
		margie.Villager,
		opal.Villager,
		paolo.Villager,
		tia.Villager,
		tucker.Villager)
)
