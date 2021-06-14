package species

// Species is an Amimal Crossing animal species.
type Species interface {
	// Class is the class of the species.
	Class() string
	// Conservation is the conservation status of the animal.
	Conservation() string
	// Domain is the domain of the species.
	Domain() string
	// Genus is the genus of the species.
	Genus() string
	// Family is the family of the species.
	Family() string
	// Fictional is whether the species is imaginary.
	Fictional() bool
	// Kingdom is the kingdom of the species.
	Kingdom() string
	// Name is the name of the species.
	Name() string
	// Order is the kingdom of the species.
	Order() string
	// Phylum is the phylum of the species.
	Phylum() string
	// Species is the species of the species.
	Species() string
}

// species implements the Species interface.
type species struct {
	class        string
	conservation string
	domain       string
	genus        string
	family       string
	fictional    bool
	kingdom      string
	name         string
	order        string
	phylum       string
	species      string
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

// Genus returns the species genus.
func (s species) Genus() string {
	return s.genus
}

// Family returns the species family.
func (s species) Family() string {
	return s.family
}

// Fictional returns the species fictional status.
func (s species) Fictional() bool {
	return s.fictional
}

// Kingdom returns the species kingdom.
func (s species) Kingdom() string {
	return s.kingdom
}

// Name returns the species name.
func (s species) Name() string {
	return s.name
}

// Order returns the species order.
func (s species) Order() string {
	return s.order
}

// Phylum returns the species phylum.
func (s species) Phylum() string {
	return s.phylum
}

// Species returns the species.
func (s species) Species() string {
	return s.species
}

var (
	// validate species implements Species.
	_ Species = species{}
)

const (
	fictional bool = true
)

const (
	na string = ""
)
