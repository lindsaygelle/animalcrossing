package duck

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/duck/bill"
	"github.com/lindsaygelle/animalcrossing/villager/duck/deena"
	"github.com/lindsaygelle/animalcrossing/villager/duck/derwin"
	"github.com/lindsaygelle/animalcrossing/villager/duck/drake"
	"github.com/lindsaygelle/animalcrossing/villager/duck/freckles"
	"github.com/lindsaygelle/animalcrossing/villager/duck/fruity"
	"github.com/lindsaygelle/animalcrossing/villager/duck/gloria"
	"github.com/lindsaygelle/animalcrossing/villager/duck/joey"
	"github.com/lindsaygelle/animalcrossing/villager/duck/ketchup"
	"github.com/lindsaygelle/animalcrossing/villager/duck/maelle"
	"github.com/lindsaygelle/animalcrossing/villager/duck/mallary"
	"github.com/lindsaygelle/animalcrossing/villager/duck/miranda"
	"github.com/lindsaygelle/animalcrossing/villager/duck/molly"
	"github.com/lindsaygelle/animalcrossing/villager/duck/pate"
	"github.com/lindsaygelle/animalcrossing/villager/duck/pompom"
	"github.com/lindsaygelle/animalcrossing/villager/duck/quillson"
	"github.com/lindsaygelle/animalcrossing/villager/duck/scoot"
	"github.com/lindsaygelle/animalcrossing/villager/duck/shinabiru"
	"github.com/lindsaygelle/animalcrossing/villager/duck/weber"
)

var (
	Villagers = villager.NewVillagers(
		bill.Villager,
		deena.Villager,
		derwin.Villager,
		drake.Villager,
		freckles.Villager,
		fruity.Villager,
		gloria.Villager,
		joey.Villager,
		ketchup.Villager,
		maelle.Villager,
		mallary.Villager,
		miranda.Villager,
		molly.Villager,
		pate.Villager,
		pompom.Villager,
		quillson.Villager,
		scoot.Villager,
		shinabiru.Villager,
		weber.Villager)
)
