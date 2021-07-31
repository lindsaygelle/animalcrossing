package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	pelicanCode    string = "pgb/plm"
	pelicanId      string = "pelican"
	pelicanUnicode string = ""
)

const (
	pelicanFictional bool = false
	pelicanSpecial   bool = false
)

var (
	// pelicanNameAmericanEnglish is the name of an Pelican in American English.
	pelicanNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(pelicanId)), 0}
)

var (
	// pelicanNames are the names of an Pelican in various languages.
	pelicanNames = names{
		language.AmericanEnglish: pelicanNameAmericanEnglish}
)

var (
	// Pelican is an Animal Crossing animal type.
	Pelican Animal = animal{
		code:      pelicanCode,
		id:        pelicanId,
		fictional: pelicanFictional,
		names:     pelicanNames,
		special:   pelicanSpecial,
		unicode:   pelicanUnicode}
)
