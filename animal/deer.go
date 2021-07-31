package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	deerCode    string = "der"
	deerId      string = "deer"
	deerUnicode string = "🦌"
)

const (
	deerFictional bool = false
	deerSpecial   bool = false
)

var (
	// deer is the name of an Deer in American English.
	deerNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(deerId)), 0}
)

var (
	// deerNames are the names of an Deer in various languages.
	deerNames = names{
		language.AmericanEnglish: deerNameAmericanEnglish}
)

var (
	// Deer is an Animal Crossing animal type.
	Deer Animal = animal{
		code:      deerCode,
		id:        deerId,
		fictional: deerFictional,
		names:     deerNames,
		special:   deerSpecial,
		unicode:   deerUnicode}
)
