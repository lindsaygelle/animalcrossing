package clay

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Clay"
	nameCanadianFrench       string = "Guido"
	nameDutch                string = "Clay"
	nameFrench               string = "Guido"
	nameGerman               string = "Dietmar"
	nameItalian              string = "Carmelo"
	nameJapanese             string = "どぐろう"
	nameLatinAmericanSpanish string = "Boliche"
	nameKorean               string = "햄둥"
	nameRussian              string = "Клэй"
	nameSpanish              string = "Boliche"
	nameSimplifiedChinese    string = "子墨"
	nameTraditionalChinese   string = "子墨"
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
