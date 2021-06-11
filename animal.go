package animalcrossing

// Animal is a type of Animal Crossing animal.
type Animal interface {
	// Name is the name of the animal.
	Name() string
}

// animal implements the Animal interface.
type animal struct {
	name string
}

// Name returns the animal's name.
func (a animal) Name() string {
	return a.name
}

var (
	// validate animal implements Animal.
	_ Animal = animal{}
)

var (
	// Animal is an array of Animal Crossing animals.
	Animals = [...]Animal{
		Alligator,
		Alpaca,
		Anteater,
		Axolotl,
		Bear,
		Beaver,
		Bird,
		Boar,
		Bovine,
		Bull,
		Camel,
		Cat,
		Chameleon,
		Chicken,
		Cow,
		Deer,
		Dodo,
		Dog,
		Duck,
		Eagle,
		Elephant,
		FrillNeckedLizard,
		FurSeal,
		Giraffe,
		Goat,
		Gorilla,
		Gull,
		Gyroid,
		Hamster,
		Hedgehog,
		Hippo,
		Kangaroo,
		Kappa,
		Koala,
		Lion,
		Mole,
		Monkey,
		Mouse,
		Octopus,
		Ostrich,
		Otter,
		Owl,
		Panther,
		Peacock,
		Pelican,
		Penguin,
		Pig,
		Pigeon}
)

