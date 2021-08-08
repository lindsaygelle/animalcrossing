package kangaroo

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/kangaroo/astrid"
	"github.com/lindsaygelle/animalcrossing/villager/kangaroo/carrie"
	"github.com/lindsaygelle/animalcrossing/villager/kangaroo/kitt"
	"github.com/lindsaygelle/animalcrossing/villager/kangaroo/koharu"
	"github.com/lindsaygelle/animalcrossing/villager/kangaroo/marcie"
	"github.com/lindsaygelle/animalcrossing/villager/kangaroo/marcy"
	"github.com/lindsaygelle/animalcrossing/villager/kangaroo/mathilda"
	"github.com/lindsaygelle/animalcrossing/villager/kangaroo/rooney"
	"github.com/lindsaygelle/animalcrossing/villager/kangaroo/sylvia"
	"github.com/lindsaygelle/animalcrossing/villager/kangaroo/valise"
	"github.com/lindsaygelle/animalcrossing/villager/kangaroo/walt"
)

var (
	Villagers = villager.NewVillagers(
		astrid.Villager,
		carrie.Villager,
		kitt.Villager,
		koharu.Villager,
		marcie.Villager,
		marcy.Villager,
		mathilda.Villager,
		rooney.Villager,
		sylvia.Villager,
		valise.Villager,
		walt.Villager)
)
