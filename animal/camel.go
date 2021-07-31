package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	camelCode    string = "cml"
	camelId      string = "camel"
	camelUnicode string = "🐪"
)

const (
	camelFictional bool = false
	camelSpecial   bool = false
)

var (
	// camel is the name of an Camel in American English.
	camelNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(camelId)), 0}
)

var (
	// camelNames are the names of an Camel in various languages.
	camelNames = names{
		language.AmericanEnglish: camelNameAmericanEnglish}
)

var (
	// Camel is an Animal Crossing animal type.
	Camel Animal = animal{
		code:      camelCode,
		id:        camelId,
		fictional: camelFictional,
		names:     camelNames,
		special:   camelSpecial,
		unicode:   camelUnicode}
)
