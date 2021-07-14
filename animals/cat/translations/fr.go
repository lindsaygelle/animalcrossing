package translations

import "github.com/lindsaygelle/animalcrossing/languages/fr"

type Fr struct {
	fr.Fr
}

func (f Fr) Value() string {
	return "Chat"
}
