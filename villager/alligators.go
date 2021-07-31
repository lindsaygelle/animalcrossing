package villager

import (
	"strings"
	"time"

	"github.com/lindsaygelle/animalcrossing/animal"
	"github.com/lindsaygelle/animalcrossing/astrology"
	"github.com/lindsaygelle/animalcrossing/birthday"
	"github.com/lindsaygelle/animalcrossing/personality"
	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

// Alfonso

const (
	alligatorAlfonsoId string = "alfonso"
)

const (
	alligatorAlfonsoSpecial bool = false
)

var (
	// alligatorAlfonsoBirthday is Alfonso's birthday.
	alligatorAlfonsoBirthday = birthday.NewBirthday(
		9, time.June)
)

var (
	// alligatorAlfonsoCatchphraseAmericanEnglish is Alfonso's catchphrase in American English.
	alligatorAlfonsoCatchphraseAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, "It's a me")
)

var (
	// alligatorAlfonsoCatchphrases is Alfonso's catchphrases in various languages.
	alligatorAlfonsoCatchphrases = translation.NewTranslations(
		alligatorAlfonsoCatchphraseAmericanEnglish)
)

var (
	// alligatorAlfonsoNameAmericanEnglish is Alfonso's name in American English.
	alligatorAlfonsoNameAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, strings.Title(alligatorAlfonsoId))
)

var (
	// alligatorAlfonsoNames is Alfonso's name in various languages.
	alligatorAlfonsoNames = translation.NewTranslations(
		alligatorAlfonsoNameAmericanEnglish)
)

var (
	// AlligatorAlfonso is an Animal Crossing villager.
	AlligatorAlfonso Villager = villager{
		animal:       animal.Alligator,
		astrology:    astrology.Gemini,
		birthday:     alligatorAlfonsoBirthday,
		catchphrases: alligatorAlfonsoCatchphrases,
		id:           alligatorAlfonsoId,
		names:        alligatorAlfonsoNames,
		personality:  personality.Lazy,
		special:      alligatorAlfonsoSpecial}
)
