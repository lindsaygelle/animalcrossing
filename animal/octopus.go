package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	octopusCode    string = "ocp"
	octopusId      string = "octopus"
	octopusUnicode string = "🐙"
)

const (
	octopusFictional bool = false
	octopusSpecial   bool = false
)

var (
	// octopusNameAmericanEnglish is the name of an Octopus in American English.
	octopusNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(octopusId)), 0}
)

var (
	// octopusNames are the names of an Octopus in various languages.
	octopusNames = names{
		language.AmericanEnglish: octopusNameAmericanEnglish}
)

var (
	// Octopus is an Animal Crossing animal type.
	Octopus Animal = animal{
		code:      octopusCode,
		id:        octopusId,
		fictional: octopusFictional,
		names:     octopusNames,
		special:   octopusSpecial,
		unicode:   octopusUnicode}
)
