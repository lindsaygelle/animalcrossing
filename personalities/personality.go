package personalities

// Personality is the personality of an Animal Crossing villager.
type Personality interface {
	// Name is the name of the personality.
	Name() string
}

// personality implements the Personality interface.
type personality struct {
	name string
}

func (p personality) Name() string {
	return p.name
}

var (
	_ Personality = personality{}
)
