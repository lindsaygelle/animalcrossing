package pig

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/pig/agnes"
	"github.com/lindsaygelle/animalcrossing/villager/pig/boris"
	"github.com/lindsaygelle/animalcrossing/villager/pig/chops"
	"github.com/lindsaygelle/animalcrossing/villager/pig/cobb"
	"github.com/lindsaygelle/animalcrossing/villager/pig/curly"
	"github.com/lindsaygelle/animalcrossing/villager/pig/gala"
	"github.com/lindsaygelle/animalcrossing/villager/pig/ganon"
	"github.com/lindsaygelle/animalcrossing/villager/pig/hambo"
	"github.com/lindsaygelle/animalcrossing/villager/pig/hugh"
	"github.com/lindsaygelle/animalcrossing/villager/pig/kevin"
	"github.com/lindsaygelle/animalcrossing/villager/pig/lucy"
	"github.com/lindsaygelle/animalcrossing/villager/pig/maggie"
	"github.com/lindsaygelle/animalcrossing/villager/pig/pancetti"
	"github.com/lindsaygelle/animalcrossing/villager/pig/peggy"
	"github.com/lindsaygelle/animalcrossing/villager/pig/pigleg"
	"github.com/lindsaygelle/animalcrossing/villager/pig/rasher"
	"github.com/lindsaygelle/animalcrossing/villager/pig/spork"
	"github.com/lindsaygelle/animalcrossing/villager/pig/suee"
	"github.com/lindsaygelle/animalcrossing/villager/pig/truffles"
)

var (
	Villagers = villager.NewVillagers(
		agnes.Villager,
		boris.Villager,
		chops.Villager,
		cobb.Villager,
		curly.Villager,
		gala.Villager,
		ganon.Villager,
		hambo.Villager,
		hugh.Villager,
		kevin.Villager,
		lucy.Villager,
		maggie.Villager,
		pancetti.Villager,
		peggy.Villager,
		pigleg.Villager,
		rasher.Villager,
		spork.Villager,
		suee.Villager,
		truffles.Villager)
)
