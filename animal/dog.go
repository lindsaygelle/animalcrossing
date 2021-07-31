package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	dogCode    string = "dog"
	dogId      string = "dog"
	dogUnicode string = "🐕"
)

const (
	dogFictional bool = false
	dogSpecial   bool = false
)

var (
	// dogNameAmericanEnglish is the name of an Dog in American English.
	dogNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(dogId)), 0}
)

var (
	// dogNames are the names of an Dog in various languages.
	dogNames = names{
		language.AmericanEnglish: dogNameAmericanEnglish}
)

var (
	// Dog is an Animal Crossing animal type.
	Dog Animal = animal{
		code:      dogCode,
		id:        dogId,
		fictional: dogFictional,
		names:     dogNames,
		special:   dogSpecial,
		unicode:   dogUnicode}
)
