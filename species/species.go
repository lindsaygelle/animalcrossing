package species

type Species interface {
	Animal() string
	Class() string
	Conservation() string
	Domain() string
	Family() string
	Fictional() bool
	Genus() string
	Kingdom() string
	Order() string
	Phylum() string
	Species() string
}
