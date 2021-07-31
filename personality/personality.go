package personality

// Personality is a personality type for an Animal Crossing villager.
type Personality interface {
	Id() string
	Names() Names
}

// personality implements Personality.
type personality struct {
	id    string
	names names
}

func (v personality) Id() string {
	return v.id
}

func (v personality) Names() Names {
	return v.names
}

var (
	_ Personality = personality{}
)
