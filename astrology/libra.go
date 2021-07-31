package astrology

import (
	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	libraId      string = "libra"
	libraSymbol  string = "♎︎"
	libraUnicode string = "♎"
)

const (
	libraLongitude uint16 = 180
)

var (
	// libraAmericanEnglish is the name of Libra in American English.
	libraAmericanEnglish = name{
		translation.New(language.AmericanEnglish, libraId)}
)

var (
	// libraNames are the names of Libra in various languages.
	libraNames = names{
		language.AmericanEnglish: libraAmericanEnglish}
)

var (
	// Libra is the astrological information for the star sign Libra.
	Libra Astrology = astrology{
		id:         libraId,
		longtitude: libraLongitude,
		names:      libraNames,
		symbol:     libraSymbol,
		unicode:    libraUnicode}
)
