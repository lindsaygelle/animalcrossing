package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	turtleCode    string = "kpg/wip/kpp/kps/kpm"
	turtleId      string = "turtle"
	turtleUnicode string = "🐢"
)

const (
	turtleFictional bool = false
	turtleSpecial   bool = false
)

var (
	// turtle is the name of an Turtle in American English.
	turtleNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(turtleId)), 0}
)

var (
	// turtleNames are the names of an Turtle in various languages.
	turtleNames = names{
		language.AmericanEnglish: turtleNameAmericanEnglish}
)

var (
	// Turtle is an Animal Crossing animal type.
	Turtle Animal = animal{
		code:      turtleCode,
		id:        turtleId,
		fictional: turtleFictional,
		names:     turtleNames,
		special:   turtleSpecial,
		unicode:   turtleUnicode}
)
