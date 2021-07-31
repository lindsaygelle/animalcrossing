package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	frogCode    string = "flg"
	frogId      string = "frog"
	frogUnicode string = "🐸"
)

const (
	frogFictional bool = false
	frogSpecial   bool = false
)

var (
	// frog is the name of an Frog in American English.
	frogNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(frogId)), 0}
)

var (
	// frogNames are the names of an Frog in various languages.
	frogNames = names{
		language.AmericanEnglish: frogNameAmericanEnglish}
)

var (
	// Frog is an Animal Crossing animal type.
	Frog Animal = animal{
		code:      frogCode,
		id:        frogId,
		fictional: frogFictional,
		names:     frogNames,
		special:   frogSpecial,
		unicode:   frogUnicode}
)
