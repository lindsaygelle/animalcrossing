package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	goatCode    string = "goa"
	goatId      string = "goat"
	goatUnicode string = "🐐"
)

const (
	goatFictional bool = false
	goatSpecial   bool = false
)

var (
	// goat is the name of an Goat in American English.
	goatNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(goatId)), 0}
)

var (
	// goatNames are the names of an Goat in various languages.
	goatNames = names{
		language.AmericanEnglish: goatNameAmericanEnglish}
)

var (
	// Goat is an Animal Crossing animal type.
	Goat Animal = animal{
		code:      goatCode,
		id:        goatId,
		fictional: goatFictional,
		names:     goatNames,
		special:   goatSpecial,
		unicode:   goatUnicode}
)
