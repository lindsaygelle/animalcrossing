package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	eagleCode    string = "pbr"
	eagleId      string = "eagle"
	eagleUnicode string = "🦅"
)

const (
	eagleFictional bool = false
	eagleSpecial   bool = false
)

var (
	// eagleNameAmericanEnglish is the name of an Eagle in American English.
	eagleNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(eagleId)), 0}
)

var (
	// eagleNames are the names of an Eagle in various languages.
	eagleNames = names{
		language.AmericanEnglish: eagleNameAmericanEnglish}
)

var (
	// Eagle is an Animal Crossing animal type.
	Eagle Animal = animal{
		code:      eagleCode,
		id:        eagleId,
		fictional: eagleFictional,
		names:     eagleNames,
		special:   eagleSpecial,
		unicode:   eagleUnicode}
)
