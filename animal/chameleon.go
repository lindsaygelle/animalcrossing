package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	chameleonCode    string = "chm"
	chameleonId      string = "chameleon"
	chameleonUnicode string = "🦎"
)

const (
	chameleonFictional bool = false
	chameleonSpecial   bool = false
)

var (
	// chameleonNameAmericanEnglish is the name of an Chameleon in American English.
	chameleonNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(chameleonId)), 0}
)

var (
	// chameleonNames are the names of an Chameleon in various languages.
	chameleonNames = names{
		language.AmericanEnglish: chameleonNameAmericanEnglish}
)

var (
	// Chameleon is an Animal Crossing animal type.
	Chameleon Animal = animal{
		code:      chameleonCode,
		id:        chameleonId,
		fictional: chameleonFictional,
		names:     chameleonNames,
		special:   chameleonSpecial,
		unicode:   chameleonUnicode}
)
