package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	pantherCode    string = "bpt"
	pantherId      string = "panther"
	pantherUnicode string = "🐈‍⬛"
)

const (
	pantherFictional bool = false
	pantherSpecial   bool = false
)

var (
	// pantherNameAmericanEnglish is the name of an Panther in American English.
	pantherNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(pantherId)), 0}
)

var (
	// pantherNames are the names of an Panther in various languages.
	pantherNames = names{
		language.AmericanEnglish: pantherNameAmericanEnglish}
)

var (
	// Panther is an Animal Crossing animal type.
	Panther Animal = animal{
		code:      pantherCode,
		id:        pantherId,
		fictional: pantherFictional,
		names:     pantherNames,
		special:   pantherSpecial,
		unicode:   pantherUnicode}
)
