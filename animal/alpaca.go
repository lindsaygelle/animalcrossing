package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	alpacaCode    string = "alp/alw"
	alpacaId      string = "alpaca"
	alpacaUnicode string = "🦙"
)

const (
	alpacaFictional bool = false
	alpacaSpecial   bool = false
)

var (
	// alpaca is the name of an Alpaca in American English.
	alpacaNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(alpacaId)), 0}
)

var (
	// alpacaNames are the names of an Alpaca in various languages.
	alpacaNames = names{
		language.AmericanEnglish: alpacaNameAmericanEnglish}
)

var (
	// Alpaca is an Animal Crossing animal type.
	Alpaca Animal = animal{
		code:      alpacaCode,
		id:        alpacaId,
		fictional: alpacaFictional,
		names:     alpacaNames,
		special:   alpacaSpecial,
		unicode:   alpacaUnicode}
)
