package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	boarCode    string = "boc/boa"
	boarId      string = "boar"
	boarUnicode string = "🐗"
)

const (
	boarFictional bool = false
	boarSpecial   bool = false
)

var (
	// boarNameAmericanEnglish is the name of an Boar in American English.
	boarNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(boarId)), 0}
)

var (
	// boarNames are the names of an Boar in various languages.
	boarNames = names{
		language.AmericanEnglish: boarNameAmericanEnglish}
)

var (
	// Boar is an Animal Crossing animal type.
	Boar Animal = animal{
		code:      boarCode,
		id:        boarId,
		fictional: boarFictional,
		names:     boarNames,
		special:   boarSpecial,
		unicode:   boarUnicode}
)
