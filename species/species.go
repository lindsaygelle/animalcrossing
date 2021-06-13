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
	// Kingdom is the kingdom of the species.
	Kingdom() string
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
	kingdom      string
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

// Species returns the species.
func (s species) Species() string {
	return s.species
}

var (
	// validate species implements Species.
	_ Species = species{}
)

// class
const (
	amphibia string = "Amphibia"
	aves     string = "Aves"
	mammalia string = "Mammalia"
	reptilia string = "Reptilia"
)

// conservation
const (
	criticallyEndangered string = "Critically Endangered"
	domesticated         string = "Domesticated"
	leastConcern         string = "Least Concern"
	vulnerable           string = "Vulnerable"
)

// domain
const (
	eukarya string = "Eukarya"
)

// family
const (
	alligatoridae  string = "Alligatoridae"
	ambystomatidae string = "Ambystomatidae"
	bovidae        string = "Bovidae"
	camelidae      string = "Camelidae"
	castoridae     string = "Castoridae"
	chamaeleonidae string = "Chamaeleonidae"
	felidae        string = "Felidae"
	phasianidae    string = "Phasianidae"
	suidae         string = "Suidae"
	ursidae        string = "Ursidae"
)

// genus
const (
	alligator string = "Alligator"
	ambystoma string = "Ambystoma"
	bos       string = "Bos"
	camelus   string = "Camelus"
	castor    string = "Castor"
	felis     string = "Felis"
	gallus    string = "Gallus"
	sus       string = "Sus"
	vicugna   string = "Vicugna"
)

// kindgom
const (
	animalia string = "Animalia"
)

// order
const (
	artiodactyla string = "Artiodactyla"
	carnivora    string = "Carnivora"
	caudata      string = "Caudata"
	crocodylia   string = "Crocodylia"
	galliformes  string = "Galliformes"
	pilosa       string = "Pilosa"
	rodentia     string = "Rodentia"
	squamata     string = "Squamata"
)

// phylum
const (
	chordata string = "Chordata"
)

// species
const (
	aMexicanum string = "A. mexicanum"
	bTaurus    string = "B. taurus"
	fCatus     string = "F. catus"
	gGallus    string = "G. gallus"
	sScrofa    string = "S. scrofa"
)
