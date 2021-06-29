package villagers

// Villager is an Animal Crossing villager.
type Villager interface {
	// Animal is the species of villager.
	Animal() string
	// Name is the name of the villager.
	Name() string
}
