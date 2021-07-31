package astrology

import (
	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	virgoId      string = "virgo"
	virgoSymbol  string = "♍︎"
	virgoUnicode string = "♍"
)

const (
	virgoLongitude uint16 = 150
)

var (
	virgoAmericanEnglish = name{
		translation.New(language.AmericanEnglish, virgoId)}
)

var (
	virgoNames = names{
		language.AmericanEnglish: virgoAmericanEnglish}
)

var (
	// Virgo is the astrological information for the star sign Virgo.
	Virgo Astrology = astrology{
		id:         virgoId,
		longtitude: virgoLongitude,
		names:      virgoNames,
		symbol:     virgoSymbol,
		unicode:    virgoUnicode}
)
