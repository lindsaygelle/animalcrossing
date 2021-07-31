package astrology

import (
	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	capricornusId      string = "capricornus"
	capricornusSymbol  string = "♑︎"
	capricornusUnicode string = "♑"
)

const (
	capricornusLongitude uint16 = 270
)

var (
	// capricornusAmericanEnglish is the name of Capricornus in American English.
	capricornusAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, capricornusId)
)

var (
	// capricornusNames are the names of Capricornus in various languages.
	capricornusNames = translation.NewTranslations(
		capricornusAmericanEnglish)
)

var (
	// Capricornus is the astrological information for the star sign Capricornus.
	Capricornus Astrology = astrology{
		id:         capricornusId,
		longtitude: capricornusLongitude,
		names:      capricornusNames,
		symbol:     capricornusSymbol,
		unicode:    capricornusUnicode}
)
