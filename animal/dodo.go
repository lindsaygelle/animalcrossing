package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	dodoCode    string = "doc/dod"
	dodoId      string = "dodo"
	dodoUnicode string = "🦤"
)

const (
	dodoFictional bool = false
	dodoSpecial   bool = false
)

var (
	// dodoNameAmericanEnglish is the name of an Dodo in American English.
	dodoNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(dodoId)), 0}
)

var (
	// dodoNames are the names of an Dodo in various languages.
	dodoNames = names{
		language.AmericanEnglish: dodoNameAmericanEnglish}
)

var (
	// Dodo is an Animal Crossing animal type.
	Dodo Animal = animal{
		code:      dodoCode,
		id:        dodoId,
		fictional: dodoFictional,
		names:     dodoNames,
		special:   dodoSpecial,
		unicode:   dodoUnicode}
)
