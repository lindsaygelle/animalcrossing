package squirrel

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/agents"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/blaire"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/cally"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/caroline"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/cece"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/filbert"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/hazel"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/kit"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/marshal"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/mint"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/nibbles"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/peanut"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/pecan"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/poppy"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/ricky"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/sally"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/sheldon"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/static"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/sylvana"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/tasha"
	"github.com/lindsaygelle/animalcrossing/villager/squirrel/viche"
)

var (
	Villagers = villager.NewVillagers(
		agents.Villager,
		blaire.Villager,
		cally.Villager,
		caroline.Villager,
		cece.Villager,
		filbert.Villager,
		hazel.Villager,
		kit.Villager,
		marshal.Villager,
		mint.Villager,
		nibbles.Villager,
		peanut.Villager,
		pecan.Villager,
		poppy.Villager,
		ricky.Villager,
		sally.Villager,
		sheldon.Villager,
		static.Villager,
		sylvana.Villager,
		tasha.Villager,
		viche.Villager)
)
