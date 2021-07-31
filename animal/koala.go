package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	koalaCode    string = "kal"
	koalaId      string = "koala"
	koalaUnicode string = "🐨"
)

const (
	koalaFictional bool = false
	koalaSpecial   bool = false
)

var (
	// koalaNameAmericanEnglish is the name of an Koala in American English.
	koalaNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(koalaId)), 0}
)

var (
	// koalaNames are the names of an Koala in various languages.
	koalaNames = names{
		language.AmericanEnglish: koalaNameAmericanEnglish}
)

var (
	// Koala is an Animal Crossing animal type.
	Koala Animal = animal{
		code:      koalaCode,
		id:        koalaId,
		fictional: koalaFictional,
		names:     koalaNames,
		special:   koalaSpecial,
		unicode:   koalaUnicode}
)
