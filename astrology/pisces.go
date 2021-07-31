package astrology

import (
	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	piscesId      string = "pisces"
	piscesSymbol  string = "♓︎"
	piscesUnicode string = "♓"
)

const (
	piscesLongitude uint16 = 330
)

var (
	// piscesAmericanEnglish is the name of Pisces in American English.
	piscesAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, piscesId)
)

var (
	// piscesNames are the names of Pisces in various languages.
	piscesNames = translation.NewTranslations(
		piscesAmericanEnglish)
)

var (
	// Pisces is the astrological information for the star sign Pisces.
	Pisces Astrology = astrology{
		id:         piscesId,
		longtitude: piscesLongitude,
		names:      piscesNames,
		symbol:     piscesSymbol,
		unicode:    piscesUnicode}
)
