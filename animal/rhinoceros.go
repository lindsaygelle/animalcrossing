package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	rhinocerosCode    string = "rhn"
	rhinocerosId      string = "rhinoceros"
	rhinocerosUnicode string = "🦏"
)

const (
	rhinocerosFictional bool = false
	rhinocerosSpecial   bool = false
)

var (
	// rhinoceros is the name of an Rhinoceros in American English.
	rhinocerosNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(rhinocerosId)), 0}
)

var (
	// rhinocerosNames are the names of an Rhinoceros in various languages.
	rhinocerosNames = names{
		language.AmericanEnglish: rhinocerosNameAmericanEnglish}
)

var (
	// Rhinoceros is an Animal Crossing animal type.
	Rhinoceros Animal = animal{
		code:      rhinocerosCode,
		id:        rhinocerosId,
		fictional: rhinocerosFictional,
		names:     rhinocerosNames,
		special:   rhinocerosSpecial,
		unicode:   rhinocerosUnicode}
)
