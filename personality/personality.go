package personality

import "github.com/lindsaygelle/animalcrossing/translation"

// Personality is a personality type for an Animal Crossing villager.
type Personality interface {
	Id() string
	Names() translation.Translations
}

// personality implements Personality.
type personality struct {
	id    string
	names translation.Translations
}

func (v personality) Id() string {
	return v.id
}

func (v personality) Names() translation.Translations {
	return v.names
}

var (
	_ Personality = personality{}
)
