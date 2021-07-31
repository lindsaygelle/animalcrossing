package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	pumpkinCode    string = "pkn"
	pumpkinId      string = "pumpkin"
	pumpkinUnicode string = "🎃"
)

const (
	pumpkinFictional bool = false
	pumpkinSpecial   bool = false
)

var (
	// pumpkinNameAmericanEnglish is the name of an Pumpkin in American English.
	pumpkinNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(pumpkinId)), 0}
)

var (
	// pumpkinNames are the names of an Pumpkin in various languages.
	pumpkinNames = names{
		language.AmericanEnglish: pumpkinNameAmericanEnglish}
)

var (
	// Pumpkin is an Animal Crossing animal type.
	Pumpkin Animal = animal{
		code:      pumpkinCode,
		id:        pumpkinId,
		fictional: pumpkinFictional,
		names:     pumpkinNames,
		special:   pumpkinSpecial,
		unicode:   pumpkinUnicode}
)
