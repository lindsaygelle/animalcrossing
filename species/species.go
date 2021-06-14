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
	na string = ""
)

// class
const (
	aaves       string = "Aaves"
	ammalia     string = "Ammalia"
	amphibia    string = "Amphibia"
	aves        string = "Aves"
	cephalopoda string = "Cephalopoda"
	mammalia    string = "Mammalia"
	reptilia    string = "Reptilia"
	sauropsida  string = "Sauropsida"
)

// conservation
const (
	basedOnFolklore      string = "Based On Folklore"
	criticallyEndangered string = "Critically Endangered"
	domesticated         string = "Domesticated"
	endangered           string = "Endangered"
	extinct              string = "Extinct"
	leastConcern         string = "Least Concern"
	nearThreatened       string = "Near Threatened"
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
	leporidae       string = "Leporidae"
	macropodidae    string = "Macropodidae"
	mephitidae      string = "Mephitidae"
	muridae         string = "Muridae"
	mustelidae      string = "Mustelidae"
	odobenidae      string = "Odobenidae"
	otariidae       string = "Otariidae"
	pelecanidae     string = "Pelecanidae"
	phascolarctidae string = "Phascolarctidae"
	phasianidae     string = "Phasianidae"
	rhinocerotidae  string = "Rhinocerotidae"
	sciuridae       string = "Sciuridae"
	spheniscidae    string = "Spheniscidae"
	struthionidae   string = "Struthionidae"
	suidae          string = "Suidae"
	talpidae        string = "Talpidae"
	tapiridae       string = "Tapiridae"
	testudinidae    string = "Testudinidae"
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
	meleagris      string = "Meleagris"
	mus            string = "Mus"
	nyctereutes    string = "Nyctereutes"
	odobenus       string = "Odobenus"
	ovis           string = "Ovis"
	panthera       string = "Panthera"
	pavo           string = "Pavo"
	pelecanus      string = "Pelecanus"
	phascolarctos  string = "Phascolarctos"
	raphus         string = "Raphus"
	rangifer       string = "Rangifer"
	struthio       string = "Struthio"
	sus            string = "Sus"
	tapirus        string = "Tapirus"
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
	accipitriformes  string = "Accipitriformes"
	anseriformes     string = "Anseriformes"
	anura            string = "Anura"
	artiodactyla     string = "Artiodactyla"
	calumbiformes    string = "Calumbiformes"
	carnivora        string = "Carnivora"
	caudata          string = "Caudata"
	charadriiformes  string = "Charadriiformes"
	columbiformes    string = "Columbiformes"
	crocodylia       string = "Crocodylia"
	diprotodontia    string = "Diprotodontia"
	eulipotyphla     string = "Eulipotyphla"
	galliformes      string = "Galliformes"
	lagomorpha       string = "Lagomorpha"
	octopoda         string = "Octopoda"
	pelecaniformes   string = "Pelecaniformes"
	perissodactyla   string = "Perissodactyla"
	pilosa           string = "Pilosa"
	primates         string = "Primates"
	proboscidea      string = "Proboscidea"
	rodentia         string = "Rodentia"
	sphenisciformes  string = "Sphenisciformes"
	squamata         string = "Squamata"
	strigiformes     string = "Strigiformes"
	struthioniformes string = "Struthioniformes"
	testudines       string = "Testudines"
)

// phylum
const (
	chordata string = "Chordata"
	mollusca string = "Mollusca"
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
	nProcyonoides        string = "N. procyonoides"
	oAries               string = "O. aries"
	oRosmarus            string = "O. rosmarus"
	pCinereus            string = "P. cinereus"
	pLeo                 string = "P. leo"
	rCucullatus          string = "R. cucullatus"
	rTarandus            string = "R. tarandus"
	sCamelus             string = "S. camelus"
	sScrofa              string = "S. scrofa"
	pTigris              string = "P. tigris"
	vVulpes              string = "V. vulpes"
)
