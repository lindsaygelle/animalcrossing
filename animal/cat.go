package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	catCode    string = "cat"
	catId      string = "cat"
	catUnicode string = "🐈"
)

const (
	catFictional bool = false
	catSpecial   bool = false
)

var (
	// catNameAmericanEnglish is the name of an Cat in American English.
	catNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(catId)), 0}
)

var (
	// catNames are the names of an Cat in various languages.
	catNames = names{
		language.AmericanEnglish: catNameAmericanEnglish}
)

var (
	// Cat is an Animal Crossing animal type.
	Cat Animal = animal{
		code:      catCode,
		id:        catId,
		fictional: catFictional,
		names:     catNames,
		special:   catSpecial,
		unicode:   catUnicode}
)
