package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	fursealCode    string = "bln/fsl"
	fursealId      string = "furseal"
	fursealUnicode string = "🦭"
)

const (
	fursealFictional bool = false
	fursealSpecial   bool = false
)

var (
	// furseal is the name of an Furseal in American English.
	fursealNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(fursealId)), 0}
)

var (
	// fursealNames are the names of an Furseal in various languages.
	fursealNames = names{
		language.AmericanEnglish: fursealNameAmericanEnglish}
)

var (
	// Furseal is an Animal Crossing animal type.
	Furseal Animal = animal{
		code:      fursealCode,
		id:        fursealId,
		fictional: fursealFictional,
		names:     fursealNames,
		special:   fursealSpecial,
		unicode:   fursealUnicode}
)
