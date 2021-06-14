package villagers

// Villager is an Animal Crossing villager.
type Villager interface {
	// Animal is the species of villager.
	Animal() string
	// Name is the name of the villager.
	Name() string
}

// villager implements Villager.
type villager struct {
	animal string
	name   string
}

// Animal returns the species of the villager.
func (v villager) Animal() string {
	return v.animal
}

// Name returns the name of the villager.
func (v villager) Name() string {
	return v.name
}

var (
	// validate villager implements Villager.
	_ Villager = villager{}
)
