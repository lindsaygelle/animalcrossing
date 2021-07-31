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
	piscesAmericanEnglish = name{
		translation.New(language.AmericanEnglish, piscesId)}
)

var (
	piscesNames = names{
		language.AmericanEnglish: piscesAmericanEnglish}
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
