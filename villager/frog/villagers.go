package frog

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/frog/camofrog"
	"github.com/lindsaygelle/animalcrossing/villager/frog/cousteau"
	"github.com/lindsaygelle/animalcrossing/villager/frog/croque"
	"github.com/lindsaygelle/animalcrossing/villager/frog/diva"
	"github.com/lindsaygelle/animalcrossing/villager/frog/drift"
	"github.com/lindsaygelle/animalcrossing/villager/frog/emerald"
	"github.com/lindsaygelle/animalcrossing/villager/frog/frobert"
	"github.com/lindsaygelle/animalcrossing/villager/frog/gigi"
	"github.com/lindsaygelle/animalcrossing/villager/frog/henry"
	"github.com/lindsaygelle/animalcrossing/villager/frog/huck"
	"github.com/lindsaygelle/animalcrossing/villager/frog/jambette"
	"github.com/lindsaygelle/animalcrossing/villager/frog/jeremiah"
	"github.com/lindsaygelle/animalcrossing/villager/frog/lily"
	"github.com/lindsaygelle/animalcrossing/villager/frog/prince"
	"github.com/lindsaygelle/animalcrossing/villager/frog/puddles"
	"github.com/lindsaygelle/animalcrossing/villager/frog/raddle"
	"github.com/lindsaygelle/animalcrossing/villager/frog/ribbot"
	"github.com/lindsaygelle/animalcrossing/villager/frog/sunny"
	"github.com/lindsaygelle/animalcrossing/villager/frog/tad"
	"github.com/lindsaygelle/animalcrossing/villager/frog/wartjr"
)

var (
	Villagers = villager.NewVillagers(
		camofrog.Villager,
		cousteau.Villager,
		croque.Villager,
		diva.Villager,
		drift.Villager,
		emerald.Villager,
		frobert.Villager,
		gigi.Villager,
		henry.Villager,
		huck.Villager,
		jambette.Villager,
		jeremiah.Villager,
		lily.Villager,
		prince.Villager,
		puddles.Villager,
		raddle.Villager,
		ribbot.Villager,
		sunny.Villager,
		tad.Villager,
		wartjr.Villager)
)
