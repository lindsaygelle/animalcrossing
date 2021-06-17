package animals

// Animal is a type of Animal Crossing animal.
type Animal interface {
	// Name is the name of the animal.
	Name() string
}

// animal implements the Animal interface.
type animal struct {
	name string
}

// Name returns the animals name.
func (a animal) Name() string {
	return a.name
}

var (
	// validate animal implements Animal.
	_ Animal = animal{}
)
