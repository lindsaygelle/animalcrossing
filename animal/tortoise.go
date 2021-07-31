package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	tortoiseCode    string = "dnk/ttlA"
	tortoiseId      string = "tortoise"
	tortoiseUnicode string = "🐢"
)

const (
	tortoiseFictional bool = false
	tortoiseSpecial   bool = false
)

var (
	// tortoiseNameAmericanEnglish is the name of an Tortoise in American English.
	tortoiseNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(tortoiseId)), 0}
)

var (
	// tortoiseNames are the names of an Tortoise in various languages.
	tortoiseNames = names{
		language.AmericanEnglish: tortoiseNameAmericanEnglish}
)

var (
	// Tortoise is an Animal Crossing animal type.
	Tortoise Animal = animal{
		code:      tortoiseCode,
		id:        tortoiseId,
		fictional: tortoiseFictional,
		names:     tortoiseNames,
		special:   tortoiseSpecial,
		unicode:   tortoiseUnicode}
)
