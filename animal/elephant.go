package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	elephantCode    string = "elp"
	elephantId      string = "elephant"
	elephantUnicode string = "🐘"
)

const (
	elephantFictional bool = false
	elephantSpecial   bool = false
)

var (
	// elephantNameAmericanEnglish is the name of an Elephant in American English.
	elephantNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(elephantId)), 0}
)

var (
	// elephantNames are the names of an Elephant in various languages.
	elephantNames = names{
		language.AmericanEnglish: elephantNameAmericanEnglish}
)

var (
	// Elephant is an Animal Crossing animal type.
	Elephant Animal = animal{
		code:      elephantCode,
		id:        elephantId,
		fictional: elephantFictional,
		names:     elephantNames,
		special:   elephantSpecial,
		unicode:   elephantUnicode}
)
