package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	alligatorCode    string = "crd"
	alligatorId      string = "alligator"
	alligatorUnicode string = "🐊"
)

const (
	alligatorFictional bool = false
	alligatorSpecial   bool = false
)

var (
	// alligator is the name of an Alligator in American English.
	alligatorNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(alligatorId)), 0}
)

var (
	// alligatorNames are the names of an Alligator in various languages.
	alligatorNames = names{
		language.AmericanEnglish: alligatorNameAmericanEnglish}
)

var (
	// Alligator is an Animal Crossing animal type.
	Alligator Animal = animal{
		code:      alligatorCode,
		id:        alligatorId,
		fictional: alligatorFictional,
		names:     alligatorNames,
		special:   alligatorSpecial,
		unicode:   alligatorUnicode}
)
