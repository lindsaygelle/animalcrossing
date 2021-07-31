package astrology

import (
	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	ariesId      string = "aries"
	ariesSymbol  string = "♈︎"
	ariesUnicode string = "♈"
)

const (
	ariesLongitude uint16 = 0
)

var (
	ariesAmericanEnglish = name{
		translation.New(language.AmericanEnglish, ariesId)}
)

var (
	ariesNames = names{
		language.AmericanEnglish: ariesAmericanEnglish}
)

var (
	// Aries is the astrological information for the star sign Aries.
	Aries Astrology = astrology{
		id:         ariesId,
		longtitude: ariesLongitude,
		names:      ariesNames,
		symbol:     ariesSymbol,
		unicode:    ariesUnicode}
)
