package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	tapirCode    string = "tap"
	tapirId      string = "tapir"
	tapirUnicode string = ""
)

const (
	tapirFictional bool = false
	tapirSpecial   bool = false
)

var (
	// tapirNameAmericanEnglish is the name of an Tapir in American English.
	tapirNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(tapirId)), 0}
)

var (
	// tapirNames are the names of an Tapir in various languages.
	tapirNames = names{
		language.AmericanEnglish: tapirNameAmericanEnglish}
)

var (
	// Tapir is an Animal Crossing animal type.
	Tapir Animal = animal{
		code:      tapirCode,
		id:        tapirId,
		fictional: tapirFictional,
		names:     tapirNames,
		special:   tapirSpecial,
		unicode:   tapirUnicode}
)
