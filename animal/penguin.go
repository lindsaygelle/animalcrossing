package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	penguinCode    string = "pgn"
	penguinId      string = "penguin"
	penguinUnicode string = "🐧"
)

const (
	penguinFictional bool = false
	penguinSpecial   bool = false
)

var (
	// penguinNameAmericanEnglish is the name of an Penguin in American English.
	penguinNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(penguinId)), 0}
)

var (
	// penguinNames are the names of an Penguin in various languages.
	penguinNames = names{
		language.AmericanEnglish: penguinNameAmericanEnglish}
)

var (
	// Penguin is an Animal Crossing animal type.
	Penguin Animal = animal{
		code:      penguinCode,
		id:        penguinId,
		fictional: penguinFictional,
		names:     penguinNames,
		special:   penguinSpecial,
		unicode:   penguinUnicode}
)
