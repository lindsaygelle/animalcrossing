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

func (a animal) Name() string {
	return a.name
}

var (
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
		Gorilla}
)
