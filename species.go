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

// Class returns the species class.
func (s species) Class() string {
	return s.class
}

// Conservation returns the species conservation status.
func (s species) Conservation() string {
	return s.conservation
}

// Domain returns the species domain.
func (s species) Domain() string {
	return s.domain
}

// Family returns the species family.
func (s species) Family() string {
	return s.family
}

// Kingdom returns the species kingdom.
func (s species) Kingdom() string {
	return s.kingdom
}

// Order returns the species order.
func (s species) Order() string {
	return s.order
}

// Phylum returns the species phylum.
func (s species) Phylum() string {
	return s.phylum
}

var (
	// validate species implements Species.
	_ Species = species{}
)
