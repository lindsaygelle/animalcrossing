package animalcrossing

// Species is an Amimal Crossing animal species.
type Species interface {
	// Class is the class of the species.
	Class() string
	// Conservation is the conservation status of the animal.
	Conservation() string
	// Domain is the domain of the species.
	Domain() string
	// Family is the family of the species.
	Family() string
	// Kingdom is the kingdom of the species.
	Kingdom() string
	// Order is the kingdom of the species.
	Order() string
	// Phylum is the phylum of the species.
	Phylum() string
}

// species implements the Species interface.
type species struct {
	class        string
	conservation string
	domain       string
	family       string
	kingdom      string
	order        string
	phylum       string
}

func (s species) Class() string {
	return s.class
}

func (s species) Conservation() string {
	return s.conservation
}

func (s species) Domain() string {
	return s.domain
}

func (s species) Family() string {
	return s.family
}

func (s species) Kingdom() string {
	return s.kingdom
}

func (s species) Order() string {
	return s.order
}

func (s species) Phylum() string {
	return s.phylum
}

var (
	_ Species = species{}
)
