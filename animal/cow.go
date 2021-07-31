package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	cowCode    string = "cow"
	cowId      string = "cow"
	cowUnicode string = "🐮"
)

const (
	cowFictional bool = false
	cowSpecial   bool = false
)

var (
	// cow is the name of an Cow in American English.
	cowNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(cowId)), 0}
)

var (
	// cowNames are the names of an Cow in various languages.
	cowNames = names{
		language.AmericanEnglish: cowNameAmericanEnglish}
)

var (
	// Cow is an Animal Crossing animal type.
	Cow Animal = animal{
		code:      cowCode,
		id:        cowId,
		fictional: cowFictional,
		names:     cowNames,
		special:   cowSpecial,
		unicode:   cowUnicode}
)
