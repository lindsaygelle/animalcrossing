package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	bullCode    string = "bul"
	bullId      string = "bull"
	bullUnicode string = "🐮"
)

const (
	bullFictional bool = false
	bullSpecial   bool = false
)

var (
	// bull is the name of an Bull in American English.
	bullNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(bullId)), 0}
)

var (
	// bullNames are the names of an Bull in various languages.
	bullNames = names{
		language.AmericanEnglish: bullNameAmericanEnglish}
)

var (
	// Bull is an Animal Crossing animal type.
	Bull Animal = animal{
		code:      bullCode,
		id:        bullId,
		fictional: bullFictional,
		names:     bullNames,
		special:   bullSpecial,
		unicode:   bullUnicode}
)
