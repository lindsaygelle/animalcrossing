package animals_test

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/animals/alligator"
	"github.com/lindsaygelle/animalcrossing/animals/alpaca"
	"github.com/lindsaygelle/animalcrossing/animals/anteater"
	"github.com/lindsaygelle/animalcrossing/animals/axolotl"
	"github.com/lindsaygelle/animalcrossing/animals/bear"
	"github.com/lindsaygelle/animalcrossing/animals/bear_cub"
	"github.com/lindsaygelle/animalcrossing/animals/beaver"
	"github.com/lindsaygelle/animalcrossing/animals/bird"
	"github.com/lindsaygelle/animalcrossing/animals/boar"
	"github.com/lindsaygelle/animalcrossing/animals/bull"
	"github.com/lindsaygelle/animalcrossing/animals/camel"
	"github.com/lindsaygelle/animalcrossing/animals/cat"
	"github.com/lindsaygelle/animalcrossing/animals/chameleon"
	"github.com/lindsaygelle/animalcrossing/animals/chicken"
	"github.com/lindsaygelle/animalcrossing/animals/cow"
	"github.com/lindsaygelle/animalcrossing/animals/deer"
	"github.com/lindsaygelle/animalcrossing/animals/dodo"
	"github.com/lindsaygelle/animalcrossing/animals/dog"
	"github.com/lindsaygelle/animalcrossing/animals/duck"
	"github.com/lindsaygelle/animalcrossing/animals/eagle"
	"github.com/lindsaygelle/animalcrossing/animals/elephant"
	"github.com/lindsaygelle/animalcrossing/animals/fox"
	"github.com/lindsaygelle/animalcrossing/animals/frill_necked_lizard"
	"github.com/lindsaygelle/animalcrossing/animals/frog"
	"github.com/lindsaygelle/animalcrossing/animals/fur_seal"
	"github.com/lindsaygelle/animalcrossing/animals/giraffe"
	"github.com/lindsaygelle/animalcrossing/animals/goat"
	"github.com/lindsaygelle/animalcrossing/animals/gorilla"
	"github.com/lindsaygelle/animalcrossing/animals/gyroid"
	"github.com/lindsaygelle/animalcrossing/animals/hamster"
	"github.com/lindsaygelle/animalcrossing/animals/hedgehog"
	"github.com/lindsaygelle/animalcrossing/animals/hippopotamus"
	"github.com/lindsaygelle/animalcrossing/animals/horse"
	"github.com/lindsaygelle/animalcrossing/animals/human"
	"github.com/lindsaygelle/animalcrossing/animals/kangaroo"
	"github.com/lindsaygelle/animalcrossing/animals/kappa"
	"github.com/lindsaygelle/animalcrossing/animals/koala"
	"github.com/lindsaygelle/animalcrossing/animals/lion"
	"github.com/lindsaygelle/animalcrossing/animals/mole"
	"github.com/lindsaygelle/animalcrossing/animals/monkey"
	"github.com/lindsaygelle/animalcrossing/animals/mouse"
	"github.com/lindsaygelle/animalcrossing/animals/octopus"
	"github.com/lindsaygelle/animalcrossing/animals/ostrich"
	"github.com/lindsaygelle/animalcrossing/animals/otter"
	"github.com/lindsaygelle/animalcrossing/animals/owl"
	"github.com/lindsaygelle/animalcrossing/animals/panther"
	"github.com/lindsaygelle/animalcrossing/animals/peacock"
	"github.com/lindsaygelle/animalcrossing/animals/pelican"
	"github.com/lindsaygelle/animalcrossing/animals/penguin"
	"github.com/lindsaygelle/animalcrossing/animals/pig"
	"github.com/lindsaygelle/animalcrossing/animals/pigeon"
	"github.com/lindsaygelle/animalcrossing/animals/rabbit"
	"github.com/lindsaygelle/animalcrossing/animals/raccoon"
	"github.com/lindsaygelle/animalcrossing/animals/reindeer"
	"github.com/lindsaygelle/animalcrossing/animals/rhinoceros"
	"github.com/lindsaygelle/animalcrossing/animals/seagull"
	"github.com/lindsaygelle/animalcrossing/animals/sheep"
	"github.com/lindsaygelle/animalcrossing/animals/skunk"
	"github.com/lindsaygelle/animalcrossing/animals/sloth"
	"github.com/lindsaygelle/animalcrossing/animals/snow"
	"github.com/lindsaygelle/animalcrossing/animals/squirrel"
	"github.com/lindsaygelle/animalcrossing/animals/tapir"
	"github.com/lindsaygelle/animalcrossing/animals/tiger"
	"github.com/lindsaygelle/animalcrossing/animals/tortoise"
	"github.com/lindsaygelle/animalcrossing/animals/turkey"
	"github.com/lindsaygelle/animalcrossing/animals/walrus"
	"github.com/lindsaygelle/animalcrossing/animals/wolf"
)

func TestSpecialFalse(t *testing.T) {
	var (
		collection = []animals.Animal{
			alligator.Alligator{},
			anteater.Anteater{},
			bear.Bear{},
			bear_cub.BearCub{},
			bird.Bird{},
			bull.Bull{},
			cat.Cat{},
			chicken.Chicken{},
			cow.Cow{},
			deer.Deer{},
			dog.Dog{},
			duck.Duck{},
			eagle.Eagle{},
			elephant.Elephant{},
			frog.Frog{},
			goat.Goat{},
			gorilla.Gorilla{},
			hamster.Hamster{},
			hippopotamus.Hippopotamus{},
			horse.Horse{},
			kangaroo.Kangaroo{},
			koala.Koala{},
			lion.Lion{},
			monkey.Monkey{},
			mouse.Mouse{},
			octopus.Octopus{},
			ostrich.Ostrich{},
			penguin.Penguin{},
			pig.Pig{},
			rabbit.Rabbit{},
			rhinoceros.Rhinoceros{},
			sheep.Sheep{},
			squirrel.Squirrel{},
			tiger.Tiger{},
			wolf.Wolf{}}
	)
	for i := 0; i < len(collection); i++ {
		var v animals.Animal = collection[i]
		if ok := v.Special() == false; !ok {
			t.Fatalf("Animal.Special() bool != %t", false)
		}
	}
}

func TestSpecialTrue(t *testing.T) {
	var (
		collection = []animals.Animal{
			alpaca.Alpaca{},
			axolotl.Axolotl{},
			beaver.Beaver{},
			boar.Boar{},
			camel.Camel{},
			chameleon.Chameleon{},
			dodo.Dodo{},
			fox.Fox{},
			frill_necked_lizard.FrillNeckedLizard{},
			fur_seal.FurSeal{},
			giraffe.Giraffe{},
			gyroid.Gyroid{},
			hedgehog.Hedgehog{},
			human.Human{},
			kappa.Kappa{},
			mole.Mole{},
			otter.Otter{},
			owl.Owl{},
			panther.Panther{},
			peacock.Peacock{},
			pelican.Pelican{},
			pigeon.Pigeon{},
			raccoon.Raccoon{},
			reindeer.Reindeer{},
			seagull.Seagull{},
			skunk.Skunk{},
			sloth.Sloth{},
			snow.Snow{},
			tapir.Tapir{},
			tortoise.Tortoise{},
			turkey.Turkey{},
			walrus.Walrus{}}
	)
	for i := 0; i < len(collection); i++ {
		var v animals.Animal = collection[i]
		if ok := v.Special() == true; !ok {
			t.Fatalf("Animal.Special() bool != %t", true)
		}
	}
}
