package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	birdCode    string = "brd"
	birdId      string = "bird"
	birdUnicode string = "🐦"
)

const (
	birdFictional bool = false
	birdSpecial   bool = false
)

var (
	// birdNameAmericanEnglish is the name of an Bird in American English.
	birdNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(birdId)), 0}
)

var (
	// birdNames are the names of an Bird in various languages.
	birdNames = names{
		language.AmericanEnglish: birdNameAmericanEnglish}
)

var (
	// Bird is an Animal Crossing animal type.
	Bird Animal = animal{
		code:      birdCode,
		id:        birdId,
		fictional: birdFictional,
		names:     birdNames,
		special:   birdSpecial,
		unicode:   birdUnicode}
)
