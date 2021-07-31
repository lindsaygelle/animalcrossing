package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	seagullCode    string = "gulB/seg/gul"
	seagullId      string = "seagull"
	seagullUnicode string = ""
)

const (
	seagullFictional bool = false
	seagullSpecial   bool = false
)

var (
	// seagullNameAmericanEnglish is the name of an Seagull in American English.
	seagullNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(seagullId)), 0}
)

var (
	// seagullNames are the names of an Seagull in various languages.
	seagullNames = names{
		language.AmericanEnglish: seagullNameAmericanEnglish}
)

var (
	// Seagull is an Animal Crossing animal type.
	Seagull Animal = animal{
		code:      seagullCode,
		id:        seagullId,
		fictional: seagullFictional,
		names:     seagullNames,
		special:   seagullSpecial,
		unicode:   seagullUnicode}
)
