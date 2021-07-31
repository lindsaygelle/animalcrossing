package astrology

import (
	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	leoId      string = "leo"
	leoSymbol  string = "♌︎"
	leoUnicode string = "♌"
)

const (
	leoLongitude uint16 = 120
)

var (
	// leoAmericanEnglish is the name of Leo in American English.
	leoAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, leoId)
)

var (
	// leoNames are the names of Leo in various languages.
	leoNames = translation.NewTranslations(
		leoAmericanEnglish)
)

var (
	// Leo is the astrological information for the star sign Leo.
	Leo Astrology = astrology{
		id:         leoId,
		longtitude: leoLongitude,
		names:      leoNames,
		symbol:     leoSymbol,
		unicode:    leoUnicode}
)
