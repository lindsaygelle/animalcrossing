package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	owlCode    string = "owl/ows"
	owlId      string = "owl"
	owlUnicode string = "🦉"
)

const (
	owlFictional bool = false
	owlSpecial   bool = false
)

var (
	// owl is the name of an Owl in American English.
	owlNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(owlId)), 0}
)

var (
	// owlNames are the names of an Owl in various languages.
	owlNames = names{
		language.AmericanEnglish: owlNameAmericanEnglish}
)

var (
	// Owl is an Animal Crossing animal type.
	Owl Animal = animal{
		code:      owlCode,
		id:        owlId,
		fictional: owlFictional,
		names:     owlNames,
		special:   owlSpecial,
		unicode:   owlUnicode}
)
