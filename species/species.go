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
	aaves      string = "Aaves"
	amphibia   string = "Amphibia"
	aves       string = "Aves"
	mammalia   string = "Mammalia"
	reptilia   string = "Reptilia"
	sauropsida string = "Sauropsida"
)

// conservation
const (
	basedOnFolklore      string = "Based On Folklore"
	criticallyEndangered string = "Critically Endangered"
	domesticated         string = "Domesticated"
	endangered           string = "Endangered"
	extinct              string = "Extinct"
	leastConcern         string = "Least Concern"
	unknown              string = "Unknown"
	vulnerable           string = "Vulnerable"
)

// domain
const (
	eukarya string = "Eukarya"
)

// family
const (
	accipitridae    string = "Accipitridae"
	agamidae        string = "Agamidae"
	alligatoridae   string = "Alligatoridae"
	ambystomatidae  string = "Ambystomatidae"
	anatidae        string = "Anatidae"
	bovidae         string = "Bovidae"
	camelidae       string = "Camelidae"
	canidae         string = "Canidae"
	castoridae      string = "Castoridae"
	cervidae        string = "Cervidae"
	chamaeleonidae  string = "Chamaeleonidae"
	cricetidae      string = "Cricetidae"
	columbidae      string = "Columbidae"
	elephantidae    string = "Elephantidae"
	equidae         string = "Equidae"
	erinaceidae     string = "Erinaceidae"
	felidae         string = "Felidae"
	giraffidae      string = "Giraffidae"
	hippopotamidea  string = "Hippopotamidea"
	hominidae       string = "Hominidae"
	laridae         string = "Laridae"
	macropodidae    string = "Macropodidae"
	otariidae       string = "Otariidae"
	phascolarctidae string = "Phascolarctidae"
	phasianidae     string = "Phasianidae"
	suidae          string = "Suidae"
	ursidae         string = "Ursidae"
)

// genus
const (
	alligator      string = "Alligator"
	ambystoma      string = "Ambystoma"
	bos            string = "Bos"
	camelus        string = "Camelus"
	canis          string = "Canis"
	capra          string = "Capra"
	castor         string = "Castor"
	chlamydosaurus string = "Chlamydosaurus"
	equus          string = "Equus"
	felis          string = "Felis"
	gallus         string = "Gallus"
	giraffa        string = "Giraffa"
	gorilla        string = "Gorilla"
	hippopotamus   string = "Hippopotamus"
	macropus       string = "Macropus"
	phascolarctos  string = "Phascolarctos"
	raphus         string = "Raphus"
	sus            string = "Sus"
	vicugna        string = "Vicugna"
	vulpes         string = "Vulpes"
)

// kindgom
const (
	animalia string = "Animalia"
	enimalia string = "Enimalia"
)

// order
const (
	accipitriformes string = "Accipitriformes"
	anseriformes    string = "Anseriformes"
	anura           string = "Anura"
	artiodactyla    string = "Artiodactyla"
	carnivora       string = "Carnivora"
	caudata         string = "Caudata"
	charadriiformes string = "Charadriiformes"
	columbiformes   string = "Columbiformes"
	crocodylia      string = "Crocodylia"
	diprotodontia   string = "Diprotodontia"
	eulipotyphla    string = "Eulipotyphla"
	galliformes     string = "Galliformes"
	perissodactyla  string = "Perissodactyla"
	pilosa          string = "Pilosa"
	primates        string = "Primates"
	proboscidea     string = "Proboscidea"
	rodentia        string = "Rodentia"
	squamata        string = "Squamata"
)

// phylum
const (
	chordata string = "Chordata"
)

// species
const (
	aMexicanum           string = "A. mexicanum"
	bTaurus              string = "B. taurus"
	cAegagrus            string = "C. aegagrus"
	canisLupusFamiliaris string = "Canis lupus familiaris"
	cKingii              string = "C. kingii"
	fCatus               string = "F. catus"
	gCamelopardalis      string = "G. camelopardalis"
	gGallus              string = "G. gallus"
	pCinereus            string = "P. cinereus"
	rCucullatus          string = "R. cucullatus"
	sScrofa              string = "S. scrofa"
	vVulpes              string = "V. vulpes"
)
