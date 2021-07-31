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
	leoAmericanEnglish = name{
		translation.New(language.AmericanEnglish, leoId)}
)

var (
	leoNames = names{
		language.AmericanEnglish: leoAmericanEnglish}
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
