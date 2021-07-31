package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	anteaterCode    string = "ant"
	anteaterId      string = "anteater"
	anteaterUnicode string = ""
)

const (
	anteaterFictional bool = false
	anteaterSpecial   bool = false
)

var (
	// anteaterNameAmericanEnglish is the name of an Anteater in American English.
	anteaterNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(anteaterId)), 0}
)

var (
	// anteaterNames are the names of an Anteater in various languages.
	anteaterNames = names{
		language.AmericanEnglish: anteaterNameAmericanEnglish}
)

var (
	// Anteater is an Animal Crossing animal type.
	Anteater Animal = animal{
		code:      anteaterCode,
		id:        anteaterId,
		fictional: anteaterFictional,
		names:     anteaterNames,
		special:   anteaterSpecial,
		unicode:   anteaterUnicode}
)
