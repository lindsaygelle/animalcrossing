package stinky

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Stinky"
	nameCanadianFrench       string = "Tupux"
	nameDutch                string = "Stinky"
	nameFrench               string = "Tupux"
	nameGerman               string = "Stefan"
	nameItalian              string = "Puzzolo"
	nameJapanese             string = "アセクサ"
	nameLatinAmericanSpanish string = "Micifuz"
	nameKorean               string = "땀띠"
	nameRussian              string = "Стинки"
	nameSpanish              string = "Micifuz"
	nameSimplifiedChinese    string = "男子汗"
	nameTraditionalChinese   string = "男子汗"
)

var (
	name = map[language.Tag]string{
		language.AmericanEnglish:      nameAmericanEnglish,
		language.CanadianFrench:       nameCanadianFrench,
		language.Dutch:                nameDutch,
		language.French:               nameFrench,
		language.German:               nameGerman,
		language.Italian:              nameItalian,
		language.Japanese:             nameJapanese,
		language.Korean:               nameKorean,
		language.LatinAmericanSpanish: nameLatinAmericanSpanish,
		language.Russian:              nameRussian,
		language.Spanish:              nameSpanish,
		language.SimplifiedChinese:    nameSimplifiedChinese,
		language.TraditionalChinese:   nameTraditionalChinese}
)
