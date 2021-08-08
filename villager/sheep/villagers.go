package sheep

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/baabara"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/cashmere"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/curlos"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/dom"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/etoile"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/eunice"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/frita"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/gen"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/muffy"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/pietro"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/stella"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/timbra"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/vesta"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/wendy"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/willow"
	"github.com/lindsaygelle/animalcrossing/villager/sheep/woolio"
)

var (
	Villagers = villager.NewVillagers(
		baabara.Villager,
		cashmere.Villager,
		curlos.Villager,
		dom.Villager,
		etoile.Villager,
		eunice.Villager,
		frita.Villager,
		gen.Villager,
		muffy.Villager,
		pietro.Villager,
		stella.Villager,
		timbra.Villager,
		vesta.Villager,
		wendy.Villager,
		willow.Villager,
		woolio.Villager)
)
