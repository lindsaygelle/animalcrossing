package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	ostrichCode    string = "ost"
	ostrichId      string = "ostrich"
	ostrichUnicode string = "🐦"
)

const (
	ostrichFictional bool = false
	ostrichSpecial   bool = false
)

var (
	// ostrichNameAmericanEnglish is the name of an Ostrich in American English.
	ostrichNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(ostrichId)), 0}
)

var (
	// ostrichNames are the names of an Ostrich in various languages.
	ostrichNames = names{
		language.AmericanEnglish: ostrichNameAmericanEnglish}
)

var (
	// Ostrich is an Animal Crossing animal type.
	Ostrich Animal = animal{
		code:      ostrichCode,
		id:        ostrichId,
		fictional: ostrichFictional,
		names:     ostrichNames,
		special:   ostrichSpecial,
		unicode:   ostrichUnicode}
)
