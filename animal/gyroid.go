package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	gyroidCode    string = ""
	gyroidId      string = "gyroid"
	gyroidUnicode string = ""
)

const (
	gyroidFictional bool = false
	gyroidSpecial   bool = false
)

var (
	// gyroidNameAmericanEnglish is the name of an Gyroid in American English.
	gyroidNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(gyroidId)), 0}
)

var (
	// gyroidNames are the names of an Gyroid in various languages.
	gyroidNames = names{
		language.AmericanEnglish: gyroidNameAmericanEnglish}
)

var (
	// Gyroid is an Animal Crossing animal type.
	Gyroid Animal = animal{
		code:      gyroidCode,
		id:        gyroidId,
		fictional: gyroidFictional,
		names:     gyroidNames,
		special:   gyroidSpecial,
		unicode:   gyroidUnicode}
)
