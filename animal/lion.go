package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	lionCode    string = "lon"
	lionId      string = "lion"
	lionUnicode string = "🦁"
)

const (
	lionFictional bool = false
	lionSpecial   bool = false
)

var (
	// lionNameAmericanEnglish is the name of an Lion in American English.
	lionNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(lionId)), 0}
)

var (
	// lionNames are the names of an Lion in various languages.
	lionNames = names{
		language.AmericanEnglish: lionNameAmericanEnglish}
)

var (
	// Lion is an Animal Crossing animal type.
	Lion Animal = animal{
		code:      lionCode,
		id:        lionId,
		fictional: lionFictional,
		names:     lionNames,
		special:   lionSpecial,
		unicode:   lionUnicode}
)
