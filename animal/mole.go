package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	moleCode    string = "mob/mol"
	moleId      string = "mole"
	moleUnicode string = ""
)

const (
	moleFictional bool = false
	moleSpecial   bool = false
)

var (
	// moleNameAmericanEnglish is the name of an Mole in American English.
	moleNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(moleId)), 0}
)

var (
	// moleNames are the names of an Mole in various languages.
	moleNames = names{
		language.AmericanEnglish: moleNameAmericanEnglish}
)

var (
	// Mole is an Animal Crossing animal type.
	Mole Animal = animal{
		code:      moleCode,
		id:        moleId,
		fictional: moleFictional,
		names:     moleNames,
		special:   moleSpecial,
		unicode:   moleUnicode}
)
