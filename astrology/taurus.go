package astrology

import (
	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	taurusId      string = "taurus"
	taurusSymbol  string = "♉︎"
	taurusUnicode string = "♉"
)

const (
	taurusLongitude uint16 = 30
)

var (
	taurusAmericanEnglish = name{
		translation.New(language.AmericanEnglish, taurusId)}
)

var (
	taurusNames = names{
		language.AmericanEnglish: taurusAmericanEnglish}
)

var (
	// Taurus is the astrological information for the star sign Taurus.
	Taurus Astrology = astrology{
		id:         taurusId,
		longtitude: taurusLongitude,
		names:      taurusNames,
		symbol:     taurusSymbol,
		unicode:    taurusUnicode}
)
