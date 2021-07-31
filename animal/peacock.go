package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	peacockCode    string = "pck"
	peacockId      string = "peacock"
	peacockUnicode string = "🦚"
)

const (
	peacockFictional bool = false
	peacockSpecial   bool = false
)

var (
	// peacockNameAmericanEnglish is the name of an Peacock in American English.
	peacockNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(peacockId)), 0}
)

var (
	// peacockNames are the names of an Peacock in various languages.
	peacockNames = names{
		language.AmericanEnglish: peacockNameAmericanEnglish}
)

var (
	// Peacock is an Animal Crossing animal type.
	Peacock Animal = animal{
		code:      peacockCode,
		id:        peacockId,
		fictional: peacockFictional,
		names:     peacockNames,
		special:   peacockSpecial,
		unicode:   peacockUnicode}
)
