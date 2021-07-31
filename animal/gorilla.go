package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	gorillaCode    string = "gor"
	gorillaId      string = "gorilla"
	gorillaUnicode string = "🦍"
)

const (
	gorillaFictional bool = false
	gorillaSpecial   bool = false
)

var (
	// gorillaNameAmericanEnglish is the name of an Gorilla in American English.
	gorillaNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(gorillaId)), 0}
)

var (
	// gorillaNames are the names of an Gorilla in various languages.
	gorillaNames = names{
		language.AmericanEnglish: gorillaNameAmericanEnglish}
)

var (
	// Gorilla is an Animal Crossing animal type.
	Gorilla Animal = animal{
		code:      gorillaCode,
		id:        gorillaId,
		fictional: gorillaFictional,
		names:     gorillaNames,
		special:   gorillaSpecial,
		unicode:   gorillaUnicode}
)
